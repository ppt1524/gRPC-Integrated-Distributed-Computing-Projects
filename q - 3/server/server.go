package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "q-3/MyUberPb"

	"crypto/tls"
	"crypto/x509"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type server struct {
	pb.UnimplementedRideSharingServiceServer
	mu      sync.RWMutex
	rides   map[string]*Ride
	drivers map[string]*Driver
}

type Driver struct {
	ID          string
	Available   bool
	CurrentRide string
}

type Ride struct {
	ID              string
	RiderID         string
	DriverID        string
	PickupLocation  string
	Destination     string
	Status          string // "pending", "accepted", "completed", "cancelled"
	AssignedDrivers map[string]bool
	CreatedAt       time.Time
}

func newServer() *server {
	return &server{
		rides:   make(map[string]*Ride),
		drivers: make(map[string]*Driver),
	}
}

func (s *server) RequestRide(ctx context.Context, req *pb.RideRequest) (*pb.RideResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	rideID := fmt.Sprintf("ride_%d", time.Now().UnixNano())
	ride := &Ride{
		ID:              rideID,
		RiderID:         req.RiderId,
		PickupLocation:  req.PickupLocation,
		Destination:     req.Destination,
		Status:          "pending",
		AssignedDrivers: make(map[string]bool),
		CreatedAt:       time.Now(),
	}
	s.rides[rideID] = ride

	go func() {
		if assigned := s.assignDriverToRide(ride); !assigned {
			s.mu.Lock()
			ride.Status = "cancelled"
			s.mu.Unlock()
		}
	}()

	return &pb.RideResponse{
		RideId: rideID,
		Status: "pending",
	}, nil
}

func (s *server) GetRideStatus(ctx context.Context, req *pb.RideStatusRequest) (*pb.RideStatusResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ride, exists := s.rides[req.RideId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "ride not found")
	}

	return &pb.RideStatusResponse{
		RideId:   ride.ID,
		Status:   ride.Status,
		DriverId: ride.DriverID,
	}, nil
}

func (s *server) AcceptRide(ctx context.Context, req *pb.AcceptRideRequest) (*pb.AcceptRideResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ride, exists := s.rides[req.RideId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "ride not found")
	}

	driver, exists := s.drivers[req.DriverId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "driver not found")
	}

	if !driver.Available {
		return nil, status.Errorf(codes.FailedPrecondition, "driver is not available")
	}

	if ride.Status != "pending" {
		return nil, status.Errorf(codes.FailedPrecondition, "ride is not in pending state")
	}

	ride.Status = "accepted"
	ride.DriverID = req.DriverId
	driver.Available = false
	driver.CurrentRide = req.RideId

	return &pb.AcceptRideResponse{
		Status: "accepted",
	}, nil
}

func (s *server) RejectRide(ctx context.Context, req *pb.RejectRideRequest) (*pb.RejectRideResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ride, exists := s.rides[req.RideId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "ride not found")
	}

	// driver, exists := s.drivers[req.DriverId]
	// if !exists {
	//     return nil, status.Errorf(codes.NotFound, "driver not found")
	// }

	ride.AssignedDrivers[req.DriverId] = true
	go func() {
		if assigned := s.assignDriverToRide(ride); !assigned {
			s.mu.Lock()
			ride.Status = "cancelled"
			s.mu.Unlock()
		}
	}()

	return &pb.RejectRideResponse{
		Status: "rejected",
	}, nil
}

func (s *server) CompleteRide(ctx context.Context, req *pb.RideCompletionRequest) (*pb.RideCompletionResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	ride, exists := s.rides[req.RideId]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "ride not found")
	}

	if ride.DriverID != req.DriverId {
		return nil, status.Errorf(codes.PermissionDenied, "driver is not assigned to this ride")
	}

	ride.Status = "completed"
	driver := s.drivers[req.DriverId]
	driver.Available = true
	driver.CurrentRide = ""

	return &pb.RideCompletionResponse{
		Status: "completed",
	}, nil
}

func (s *server) assignDriverToRide(ride *Ride) bool {
	time.Sleep(500 * time.Millisecond) // Simulate some processing time

	s.mu.Lock()
	defer s.mu.Unlock()

	for driverID, driver := range s.drivers {
		if driver.Available && !ride.AssignedDrivers[driverID] {
			ride.AssignedDrivers[driverID] = true

			// Start a goroutine to handle timeout
			go s.handleDriverResponseTimeout(ride, driver)
			return true
		}
	}
	return false
}

func (s *server) handleDriverResponseTimeout(ride *Ride, driver *Driver) {
	time.Sleep(30 * time.Second) // Wait for driver response

	s.mu.Lock()
	defer s.mu.Unlock()

	if ride.Status == "pending" {
		// Driver didn't respond in time, try to assign to another driver
		go func() {
			if assigned := s.assignDriverToRide(ride); !assigned {
				s.mu.Lock()
				ride.Status = "cancelled"
				s.mu.Unlock()
			}
		}()
	}
}

func main() {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatalf("failed to load server certificate: %v", err)
	}

	// Load CA's certificate to validate client certificates
	caCert, err := ioutil.ReadFile("rootCA.pem")
	if err != nil {
		log.Fatalf("failed to read CA's certificate: %v", err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	// Create the TLS credentials for the server
	creds := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.RequireAndVerifyClientCert, // Enforce mTLS
		ClientCAs:    caCertPool,
	})

	port := "50051"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	lis, err := net.Listen("tcp", ":"+port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	interceptors := ChainUnaryInterceptors(AuthorizationInterceptor, LoggingInterceptor)

	grpcServer := grpc.NewServer(grpc.Creds(creds), grpc.UnaryInterceptor(interceptors))

	srv := newServer()

	// Pre-register some drivers (In practice, you'd have driver registration)
	srv.drivers["driver1"] = &Driver{ID: "driver1", Available: true}
	// srv.drivers["driver2"] = &Driver{ID: "driver2", Available: true}

	pb.RegisterRideSharingServiceServer(grpcServer, srv)

	log.Println("Server is running with TLS on port : ", port)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

// ExtractRoleFromCert extracts the role (rider/driver) from the client certificate
func ExtractRoleFromCert(ctx context.Context) (string, error) {
	peer, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("no peer found")
	}

	if len(peer.AuthInfo.(credentials.TLSInfo).State.PeerCertificates) == 0 {
		return "", fmt.Errorf("no peer certificate provided")
	}

	cert := peer.AuthInfo.(credentials.TLSInfo).State.PeerCertificates[0]

	if cert.NotAfter.Before(time.Now()) {
		return "", fmt.Errorf("certificate has expired. Expiry time was: %s", cert.NotAfter.Format(time.RFC3339))
	}

	// In practice, you would extract a field or custom extension from the cert to determine the role
	// For simplicity, we assume the subject common name (CN) contains "rider" or "driver"
	if cert.Subject.CommonName == "rider" {
		return "rider", nil
	} else if cert.Subject.CommonName == "driver" {
		return "driver", nil
	}

	return "", fmt.Errorf("invalid role in certificate")
}

// AuthorizationInterceptor ensures only riders can request rides and only drivers can accept/complete them
func AuthorizationInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	role, err := ExtractRoleFromCert(ctx)
	if err != nil {
		return nil, status.Errorf(codes.PermissionDenied, "could not determine role: %v", err)
	}

	// Restrict actions based on the client's role
	if info.FullMethod == "/RideSharingService/RequestRide" && role != "rider" {
		return nil, status.Errorf(codes.PermissionDenied, "only riders can request rides")
	} else if (info.FullMethod == "/RideSharingService/AcceptRide" || info.FullMethod == "/RideSharingService/CompleteRide") && role != "driver" {
		return nil, status.Errorf(codes.PermissionDenied, "only drivers can accept or complete rides")
	}

	// Proceed with the request
	return handler(ctx, req)
}

// LoggingInterceptor logs method names, timestamps, and client roles
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	startTime := time.Now()

	// Extract role for logging
	role, err := ExtractRoleFromCert(ctx)
	if err != nil {
		role = "unknown"
	}

	// Proceed with the request
	resp, err := handler(ctx, req)

	// Log details
	log.Printf("Method: %s | Role: %s | Time: %s | Error: %v", info.FullMethod, role, startTime.Format(time.RFC3339), err)

	return resp, err
}

// ChainUnaryInterceptors combines multiple interceptors
func ChainUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Nest interceptors in reverse order
		outerHandler := handler
		for i := len(interceptors) - 1; i >= 0; i-- {
			outerHandler = func(currentInterceptor grpc.UnaryServerInterceptor, currentHandler grpc.UnaryHandler) grpc.UnaryHandler {
				return func(ctx context.Context, req interface{}) (interface{}, error) {
					return currentInterceptor(ctx, req, info, currentHandler)
				}
			}(interceptors[i], outerHandler)
		}
		return outerHandler(ctx, req)
	}
}
