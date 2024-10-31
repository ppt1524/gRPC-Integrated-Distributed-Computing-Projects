package main

import (
    "context"
    "log"
    "google.golang.org/grpc"
    "sort"
    pb "q-2/KNNpb"
    "io"
)

// Contact a single server
func queryServerStream(client pb.KNNServiceClient, point []float32, k int32) ([]*pb.Neighbor, error) {
    req := &pb.KNNRequest{
        DataPoint: point,
        K:         k,
    }

    stream, err := client.FindKNearestNeighbors(context.Background(), req)
    if err != nil {
        return nil, err
    }

    var neighbors []*pb.Neighbor
    for {
        neighbor, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        neighbors = append(neighbors, neighbor)
    }

    return neighbors, nil
}


func main() {
    servers := []string{"localhost:50051", "localhost:50052", "localhost:50053", "localhost:50054"}  // List of server addresses
    // servers := []string{"localhost:50051"}  // List of server addresses


    point := []float32{1.0, 1.0}  // Example query point
    k := int32(3)  // Number of neighbors to find

    allNeighbors := []*pb.Neighbor{}

    // Query all servers
    for _, addr := range servers {
        conn, err := grpc.Dial(addr, grpc.WithInsecure())
        if err != nil {
            log.Fatalf("Failed to connect to server: %v", err)
        }
        defer conn.Close()

        client := pb.NewKNNServiceClient(conn)
        neighbors, err := queryServerStream(client, point, k)
        if err != nil {
            log.Fatalf("Failed to query server: %v", err)
        }

        allNeighbors = append(allNeighbors, neighbors...)
    }

    // Sort all neighbors by distance and pick the top k
    sort.Slice(allNeighbors, func(i, j int) bool {
        return allNeighbors[i].Distance < allNeighbors[j].Distance
    })
    if len(allNeighbors) > int(k) {
        allNeighbors = allNeighbors[:k]
    }

    log.Printf("Global %d nearest neighbors:", k)
    for _, neighbor := range allNeighbors {
        log.Printf("Point: %v, Distance: %v", neighbor.Point, neighbor.Distance)
    }
}
