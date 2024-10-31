package main

import (
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "log"
    "time"
    "bufio"
    "context"
    "fmt"
    "os"
    "sync/atomic"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    pb "q-3/MyUberPb"
)

type LoadBalancer struct {
    addresses []string
    next      uint32
}

func NewLoadBalancer(addresses []string) *LoadBalancer {
    return &LoadBalancer{
        addresses: addresses,
    }
}

func (lb *LoadBalancer) GetNextAddress_RoundRobin() string {
    n := atomic.AddUint32(&lb.next, 1)
    return lb.addresses[(int(n)-1)%len(lb.addresses)]
}

func (lb *LoadBalancer) GetNextAddress_PickFirst() string {
    return lb.addresses[0]
}

func main() {
    // Load client's certificate and private key (Rider)
    riderCert, err := tls.LoadX509KeyPair("rider.crt", "rider.key")
    if err != nil {
        log.Fatalf("failed to load rider's certificate: %v", err)
    }

    // Load CA's certificate to verify the server
    caCert, err := ioutil.ReadFile("rootCA.pem")
    if err != nil {
        log.Fatalf("failed to read CA's certificate: %v", err)
    }
    caCertPool := x509.NewCertPool()
    caCertPool.AppendCertsFromPEM(caCert)

    // Create the TLS credentials for the client
    creds := credentials.NewTLS(&tls.Config{
        Certificates: []tls.Certificate{riderCert},
        RootCAs:      caCertPool,
    })

    // Initialize the load balancer with multiple server addresses
    lb := NewLoadBalancer([]string{
        "localhost:50051",
        "localhost:50052",
        "localhost:50053",
    })

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Welcome to Ride-Sharing CLI!")
    for {
        fmt.Println("Choose an action: (1) Request Ride (2) Check Ride Status (3) Exit")
        fmt.Print("> ")
        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            requestRide(lb, creds, scanner)
        case "2":
            checkRideStatus(lb, creds, scanner)
        case "3":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}

func getClient(lb *LoadBalancer, creds credentials.TransportCredentials) (pb.RideSharingServiceClient, *grpc.ClientConn) {
    address := lb.GetNextAddress_PickFirst()
    // address := lb.GetNextAddress_RoundRobin()
    conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
    if err != nil {
        log.Fatalf("did not connect to %s: %v", address, err)
    }
    return pb.NewRideSharingServiceClient(conn), conn
}

func requestRide(lb *LoadBalancer, creds credentials.TransportCredentials, scanner *bufio.Scanner) {
    client, conn := getClient(lb, creds)
    defer conn.Close()

    fmt.Print("Enter Rider ID: ")
    scanner.Scan()
    riderID := scanner.Text()

    fmt.Print("Enter Pickup Location: ")
    scanner.Scan()
    pickup := scanner.Text()

    fmt.Print("Enter Destination: ")
    scanner.Scan()
    destination := scanner.Text()

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    res, err := client.RequestRide(ctx, &pb.RideRequest{
        RiderId:        riderID,
        PickupLocation: pickup,
        Destination:    destination,
    })

    if err != nil {
        log.Fatalf("Ride request failed: %v", err)
    }
    fmt.Printf("Ride requested! Ride ID: %s, Status: %s\n", res.GetRideId(), res.GetStatus())
}

func checkRideStatus(lb *LoadBalancer, creds credentials.TransportCredentials, scanner *bufio.Scanner) {
    client, conn := getClient(lb, creds)
    defer conn.Close()

    fmt.Print("Enter Rider ID: ")
    scanner.Scan()
    riderID := scanner.Text()

    fmt.Print("Enter Ride ID: ")
    scanner.Scan()
    rideID := scanner.Text()

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    res, err := client.GetRideStatus(ctx, &pb.RideStatusRequest{
        RiderId: riderID,
        RideId:  rideID,
    })
    if err != nil {
        log.Fatalf("Failed to get ride status: %v", err)
    }
    fmt.Printf("Ride ID: %s, Status: %s, Driver ID: %s\n", res.GetRideId(), res.GetStatus(), res.GetDriverId())
}