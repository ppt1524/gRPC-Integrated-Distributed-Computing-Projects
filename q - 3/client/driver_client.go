package main

import (
    "bufio"
    "context"
    "fmt"
    "log"
    "os"
    "time"
    "crypto/tls"
    "crypto/x509"
    "io/ioutil"
    "sync/atomic"

    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    pb "q-3/MyUberPb"
)

type RoundRobinLoadBalancer struct {
    addresses []string
    next      uint32
}

func NewRoundRobinLoadBalancer(addresses []string) *RoundRobinLoadBalancer {
    return &RoundRobinLoadBalancer{
        addresses: addresses,
    }
}

func (lb *RoundRobinLoadBalancer) GetNextAddress() string {
    n := atomic.AddUint32(&lb.next, 1)
    return lb.addresses[(int(n)-1)%len(lb.addresses)]
}

func main() {
    // Load driver's certificate and private key
    driverCert, err := tls.LoadX509KeyPair("driver.crt", "driver.key")
    if err != nil {
        log.Fatalf("failed to load driver's certificate: %v", err)
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
        Certificates: []tls.Certificate{driverCert},
        RootCAs:      caCertPool,
    })

    // Initialize the load balancer with multiple server addresses
    lb := NewRoundRobinLoadBalancer([]string{
        "localhost:50051",
        "localhost:50052",
        "localhost:50053",
    })

    scanner := bufio.NewScanner(os.Stdin)

    fmt.Println("Welcome, Driver!")
    for {
        fmt.Println("Choose an action: (1) Accept Ride (2) Reject Ride (3) Complete Ride (4) Exit")
        fmt.Print("> ")
        scanner.Scan()
        choice := scanner.Text()

        switch choice {
        case "1":
            acceptRide(lb, creds, scanner)
        case "2":
            rejectRide(lb, creds, scanner)
        case "3":
            completeRide(lb, creds, scanner)
        case "4":
            fmt.Println("Exiting...")
            return
        default:
            fmt.Println("Invalid choice")
        }
    }
}

func getClient(lb *RoundRobinLoadBalancer, creds credentials.TransportCredentials) (pb.RideSharingServiceClient, *grpc.ClientConn) {
    address := lb.GetNextAddress()
    conn, err := grpc.Dial(address, grpc.WithTransportCredentials(creds))
    if err != nil {
        log.Fatalf("did not connect to %s: %v", address, err)
    }
    return pb.NewRideSharingServiceClient(conn), conn
}

func acceptRide(lb *RoundRobinLoadBalancer, creds credentials.TransportCredentials, scanner *bufio.Scanner) {
    client, conn := getClient(lb, creds)
    defer conn.Close()

    fmt.Print("Enter Driver ID: ")
    scanner.Scan()
    driverID := scanner.Text()

    fmt.Print("Enter Ride ID: ")
    scanner.Scan()
    rideID := scanner.Text()

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    res, err := client.AcceptRide(ctx, &pb.AcceptRideRequest{
        DriverId: driverID,
        RideId:   rideID,
    })
    if err != nil {
        log.Fatalf("Failed to accept ride: %v", err)
    }
    fmt.Printf("Ride accepted! Status: %s\n", res.GetStatus())
}

func rejectRide(lb *RoundRobinLoadBalancer, creds credentials.TransportCredentials, scanner *bufio.Scanner) {
    client, conn := getClient(lb, creds)
    defer conn.Close()

    fmt.Print("Enter Driver ID: ")
    scanner.Scan()
    driverID := scanner.Text()

    fmt.Print("Enter Ride ID: ")
    scanner.Scan()
    rideID := scanner.Text()

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    res, err := client.RejectRide(ctx, &pb.RejectRideRequest{
        DriverId: driverID,
        RideId:   rideID,
    })
    if err != nil {
        log.Fatalf("Failed to reject ride: %v", err)
    }
    fmt.Printf("Ride rejected! Status: %s\n", res.GetStatus())
}

func completeRide(lb *RoundRobinLoadBalancer, creds credentials.TransportCredentials, scanner *bufio.Scanner) {
    client, conn := getClient(lb, creds)
    defer conn.Close()

    fmt.Print("Enter Driver ID: ")
    scanner.Scan()
    driverID := scanner.Text()

    fmt.Print("Enter Ride ID: ")
    scanner.Scan()
    rideID := scanner.Text()

    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    res, err := client.CompleteRide(ctx, &pb.RideCompletionRequest{
        DriverId: driverID,
        RideId:   rideID,
    })
    if err != nil {
        log.Fatalf("Failed to complete ride: %v", err)
    }
    fmt.Printf("Ride completed! Status: %s\n", res.GetStatus())
}