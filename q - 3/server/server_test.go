package main

import (
    "context"
    "testing"
    "time"
    
    pb "q-3/MyUberPb"
    "github.com/stretchr/testify/assert"
)

func setupTestServer() *server {
    s := newServer()
    // Pre-register some test drivers
    s.drivers["driver1"] = &Driver{ID: "driver1", Available: true}
    s.drivers["driver2"] = &Driver{ID: "driver2", Available: true}
    return s
}

func TestRequestRide(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()

    tests := []struct {
        name        string
        request     *pb.RideRequest
        expectError bool
    }{
        {
            name: "Valid ride request",
            request: &pb.RideRequest{
                RiderId:        "rider1",
                PickupLocation: "Location A",
                Destination:    "Location B",
            },
            expectError: false,
        },
        // {
        //     name: "Empty pickup location",
        //     request: &pb.RideRequest{
        //         RiderId:        "rider1",
        //         PickupLocation: "",
        //         Destination:    "Location B",
        //     },
        //     expectError: true,
        // },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            response, err := s.RequestRide(ctx, tt.request)
            
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.NotEmpty(t, response.RideId)
                assert.Equal(t, "pending", response.Status)
            }
        })
    }
}

func TestAcceptRide(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()

    // First, create a ride
    rideReq := &pb.RideRequest{
        RiderId:        "rider1",
        PickupLocation: "Location A",
        Destination:    "Location B",
    }
    rideResp, _ := s.RequestRide(ctx, rideReq)
    rideID := rideResp.RideId

    tests := []struct {
        name        string
        request     *pb.AcceptRideRequest
        expectError bool
    }{
        {
            name: "Valid ride acceptance",
            request: &pb.AcceptRideRequest{
                DriverId: "driver1",
                RideId:   rideID,
            },
            expectError: false,
        },
        {
            name: "Invalid ride ID",
            request: &pb.AcceptRideRequest{
                DriverId: "driver1",
                RideId:   "nonexistent_ride",
            },
            expectError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            response, err := s.AcceptRide(ctx, tt.request)
            
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, "accepted", response.Status)
            }
        })
    }
}

func TestGetRideStatus(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()

    // Create a ride
    rideReq := &pb.RideRequest{
        RiderId:        "rider1",
        PickupLocation: "Location A",
        Destination:    "Location B",
    }
    rideResp, _ := s.RequestRide(ctx, rideReq)
    rideID := rideResp.RideId

    tests := []struct {
        name        string
        request     *pb.RideStatusRequest
        expectError bool
        expectStatus string
    }{
        {
            name: "Check existing ride",
            request: &pb.RideStatusRequest{
                RiderId: "rider1",
                RideId:  rideID,
            },
            expectError: false,
            expectStatus: "pending",
        },
        {
            name: "Check nonexistent ride",
            request: &pb.RideStatusRequest{
                RiderId: "rider1",
                RideId:  "nonexistent_ride",
            },
            expectError: true,
            expectStatus: "",
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            response, err := s.GetRideStatus(ctx, tt.request)
            
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.expectStatus, response.Status)
            }
        })
    }
}

func TestCompleteRide(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()

    // Create and accept a ride
    rideReq := &pb.RideRequest{
        RiderId:        "rider1",
        PickupLocation: "Location A",
        Destination:    "Location B",
    }
    rideResp, _ := s.RequestRide(ctx, rideReq)
    rideID := rideResp.RideId

    acceptReq := &pb.AcceptRideRequest{
        DriverId: "driver1",
        RideId:   rideID,
    }
    s.AcceptRide(ctx, acceptReq)

    tests := []struct {
        name        string
        request     *pb.RideCompletionRequest
        expectError bool
    }{
        {
            name: "Complete valid ride",
            request: &pb.RideCompletionRequest{
                DriverId: "driver1",
                RideId:   rideID,
            },
            expectError: false,
        },
        {
            name: "Complete nonexistent ride",
            request: &pb.RideCompletionRequest{
                DriverId: "driver1",
                RideId:   "nonexistent_ride",
            },
            expectError: true,
        },
        {
            name: "Complete ride with wrong driver",
            request: &pb.RideCompletionRequest{
                DriverId: "driver2",
                RideId:   rideID,
            },
            expectError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            response, err := s.CompleteRide(ctx, tt.request)
            
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, "completed", response.Status)
            }
        })
    }
}

func TestDriverTimeout(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()

    // Create a ride
    rideReq := &pb.RideRequest{
        RiderId:        "rider1",
        PickupLocation: "Location A",
        Destination:    "Location B",
    }
    rideResp, _ := s.RequestRide(ctx, rideReq)
    rideID := rideResp.RideId

    // Wait for timeout
    time.Sleep(30 * time.Second)

    // Check ride status
    statusReq := &pb.RideStatusRequest{
        RiderId: "rider1",
        RideId:  rideID,
    }
    statusResp, err := s.GetRideStatus(ctx, statusReq)

    assert.NoError(t, err)
    // The ride should either be cancelled or reassigned
    assert.True(t, statusResp.Status == "cancelled" || statusResp.DriverId != "")
}

// func TestMultipleConcurrentRequests(t *testing.T) {
//     s := setupTestServer()
//     ctx := context.Background()

//     numRequests := 5
//     results := make(chan string, numRequests)

//     for i := 0; i < numRequests; i++ {
//         go func() {
//             rideReq := &pb.RideRequest{
//                 RiderId:        "rider1",
//                 PickupLocation: "Location A",
//                 Destination:    "Location B",
//             }
//             resp, err := s.RequestRide(ctx, rideReq)
//             if err != nil {
//                 results <- "error"
//             } else {
// 				fmt.Println(resp.RideId)
//                 results <- resp.RideId
//             }
//         }()
//     }

// 	fmt.Println(results)
//     rideIDs := make(map[string]bool)
//     for i := 0; i < numRequests; i++ {
//         rideID := <-results
//         assert.NotEqual(t, "error", rideID)
//         // Ensure we don't get duplicate ride IDs
// 		fmt.Printf("HERE: %s",  rideID)
//         assert.False(t, rideIDs[rideID])
//         rideIDs[rideID] = true
//     }
// }

func TestRejectRide(t *testing.T) {
    s := setupTestServer()
    ctx := context.Background()

    // Create a ride
    rideReq := &pb.RideRequest{
        RiderId:        "rider1",
        PickupLocation: "Location A",
        Destination:    "Location B",
    }
    rideResp, _ := s.RequestRide(ctx, rideReq)
    rideID := rideResp.RideId

    tests := []struct {
        name        string
        request     *pb.RejectRideRequest
        expectError bool
    }{
        {
            name: "Reject valid ride",
            request: &pb.RejectRideRequest{
                DriverId: "driver1",
                RideId:   rideID,
            },
            expectError: false,
        },
        {
            name: "Reject nonexistent ride",
            request: &pb.RejectRideRequest{
                DriverId: "driver1",
                RideId:   "nonexistent_ride",
            },
            expectError: true,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            response, err := s.RejectRide(ctx, tt.request)
            
            if tt.expectError {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, "rejected", response.Status)
            }
        })
    }
}