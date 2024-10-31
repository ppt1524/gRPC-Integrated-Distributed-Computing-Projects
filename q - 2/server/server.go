package main

import (
    "encoding/csv"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "log"
    "math"
    "net"
    "os"
    "sort"
    "strconv"

    pb "q-2/KNNpb" // Replace with your actual import path for the generated protobuf files
)

// Neighbor represents a nearest neighbor with the distance.
type Neighbor struct {
    Point    []float32
    Distance float32
}

// Server holds the dataset for this gRPC server.
type server struct {
    data [][]float32
    pb.UnimplementedKNNServiceServer
}

// Euclidean distance calculation between two points.
func euclideanDistance(a, b []float32) float32 {
    var sum float32
    for i := range a {
        sum += (a[i] - b[i]) * (a[i] - b[i])
    }
    return float32(math.Sqrt(float64(sum)))
}

// Load dataset from a CSV file and parse it into a 2D slice of float32.
func loadDataset(filePath string) [][]float32 {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatalf("failed to open file: %v", err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    records, err := reader.ReadAll()
    if err != nil {
        log.Fatalf("failed to read CSV: %v", err)
    }

    dataset := [][]float32{}
    for _, record := range records {
        point := []float32{}
        for _, val := range record {
            fval, err := strconv.ParseFloat(val, 32)
            if err != nil {
                log.Fatalf("failed to parse float: %v", err)
            }
            point = append(point, float32(fval))
        }
        dataset = append(dataset, point)
    }
    return dataset
}

// FindKNearestNeighbors finds the k nearest neighbors within the server's dataset and streams the result.
// Implements the FindKNearestNeighbors RPC
func (s *server) FindKNearestNeighbors(req *pb.KNNRequest, stream pb.KNNService_FindKNearestNeighborsServer) error {
    queryPoint := req.GetDataPoint()
    k := int(req.GetK())

    // Calculate distance to all points in the dataset
    neighbors := make([]*pb.Neighbor, len(s.data))
    for i, point := range s.data {
        dist := euclideanDistance(point, queryPoint)
        neighbors[i] = &pb.Neighbor{Point: point, Distance: dist}
    }

    // Sort by distance
    sort.Slice(neighbors, func(i, j int) bool {
        return neighbors[i].Distance < neighbors[j].Distance
    })

    // Stream k nearest neighbors
    for i := 0; i < k && i < len(neighbors); i++ {
        if err := stream.Send(neighbors[i]); err != nil {
            return err
        }
    }

    return nil
}


// Partition dataset into equal-sized chunks for multiple servers.
func partitionDataset(dataset [][]float32, numPartitions int) [][][]float32 {
    partitions := make([][][]float32, numPartitions)
    for i, point := range dataset {
        partitionIndex := i % numPartitions
        partitions[partitionIndex] = append(partitions[partitionIndex], point)
    }
    return partitions
}

func main() {
	fmt.Println("Arguments: ", os.Args)
    // Load dataset from CSV file.
    dataset := loadDataset("./server/knn_dataset.csv")

    // Simulate multiple partitions (for multiple servers).
    numPartitions := 4 // Adjust based on the number of servers you're using.
    partitions := partitionDataset(dataset, numPartitions)

    // Assign this server a partition (e.g., partition 0 for this server).
    partitionId := os.Args[1] // Give this as command line for each server
    partitionIndex, err := strconv.Atoi(partitionId)
    serverData := partitions[partitionIndex]

    // Set up the gRPC server.
    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051+ (partitionIndex))) // Increment port for each server.
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterKNNServiceServer(s, &server{data: serverData})

    // Register reflection service on gRPC server (useful for tools like grpcurl).
    reflection.Register(s)

    log.Printf("Server is running on port %d", 50051+partitionIndex)
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
