package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"sync"
	
	"google.golang.org/grpc"
	pb "q-1/labyrinthpb"
)

// LabyrinthServer struct to hold game state
type LabyrinthServer struct {
	pb.UnimplementedLabyrinthServer
	mu              sync.Mutex
	width, height   int32
	labyrinth       [][]rune
	playerX, playerY int32
	playerScore     int32
	playerHealth    int32
	remainingSpells int32
}

// Create the labyrinth from a file or statically
func NewLabyrinthServer() *LabyrinthServer {
	// Static labyrinth for now (can read from file if needed)
	lab := [][]rune{
		{'E', 'E', 'W', 'C', 'E'},
		{'E', 'W', 'E', 'E', 'E'},
		{'E', 'C', 'E', 'W', 'E'},
		{'W', 'E', 'C', 'E', 'E'},
		{'C', 'E', 'E', 'E', 'E'},
	}

	return &LabyrinthServer{
		width:           5,
		height:          5,
		labyrinth:       lab,
		playerX:         0,
		playerY:         0,
		playerScore:     0,
		playerHealth:    3,
		remainingSpells: 3,
	}
}

// GetLabyrinthInfo RPC
func (s *LabyrinthServer) GetLabyrinthInfo(ctx context.Context, in *pb.Empty) (*pb.LabyrinthInfo, error) {
	labyrinthStr := make([]string, len(s.labyrinth))
    for i, row := range s.labyrinth {
		labyrinthStr[i] = ""
        for _, tile := range row {
			labyrinthStr[i] += string(tile) + " "
        }
        labyrinthStr[i] = labyrinthStr[i][:len(labyrinthStr[i])-1] // Trim trailing space
    }

	return &pb.LabyrinthInfo{
		Width:  s.width,
		Height: s.height,
		Labyrinth: labyrinthStr,
	}, nil
}

// GetPlayerStatus RPC
func (s *LabyrinthServer) GetPlayerStatus(ctx context.Context, in *pb.Empty) (*pb.PlayerStatus, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	return &pb.PlayerStatus{
		Score:          s.playerScore,
		Health:         s.playerHealth,
		X:              s.playerX,
		Y:              s.playerY,
		RemainingSpells: s.remainingSpells,
	}, nil
}

// RegisterMove RPC
func (s *LabyrinthServer) RegisterMove(ctx context.Context, req *pb.MoveRequest) (*pb.MoveResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	var newX, newY int32 = s.playerX, s.playerY

	switch req.Direction {
	case "up":
		newY--
	case "down":
		newY++
	case "left":
		newX--
	case "right":
		newX++
	}

	// Validate move
	if newX < 0 || newX >= s.width || newY < 0 || newY >= s.height {
		s.playerHealth--
		return &pb.MoveResponse{Status: "invalid move"}, nil
	}

	var isCoin string = ""
	tile := s.labyrinth[newY][newX]
	// Update player position
	s.playerX, s.playerY = newX, newY		
	switch tile {
	case 'W':
		s.playerHealth--
		if s.playerHealth <= 0 {
			return &pb.MoveResponse{Status: "Death & Game Over"}, nil
		} else if newX == s.width-1 && newY == s.height-1 {
			return &pb.MoveResponse{Status: "hit wall & victory"}, nil
		}
		s.labyrinth[newY][newX] = 'E'
		return &pb.MoveResponse{Status: "hit wall"}, nil
	case 'C':
		s.playerScore++
		s.labyrinth[newY][newX] = 'E' // Coin collected, tile becomes empty
		isCoin = " & Coin collected"
	}


	if newX == s.width-1 && newY == s.height-1 {
		return &pb.MoveResponse{Status: "victory" + isCoin}, nil
	}

	return &pb.MoveResponse{Status: "success" + isCoin}, nil
}

// Revelio RPC
func (s *LabyrinthServer) Revelio(req *pb.RevelioRequest, stream pb.Labyrinth_RevelioServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.remainingSpells--
	for y := req.Y - 1; y <= req.Y+1; y++ {
		for x := req.X - 1; x <= req.X+1; x++ {
			if x >= 0 && x < s.width && y >= 0 && y < s.height && s.labyrinth[y][x] == rune(req.TileType[0]) {
				stream.Send(&pb.TilePosition{
					X: x,
					Y: y,
				})
			}
		}
	}
	return nil
}

// Bombarda RPC
func (s *LabyrinthServer) Bombarda(stream pb.Labyrinth_BombardaServer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.remainingSpells--
	for {
		tilePos, err := stream.Recv()
		if err != nil {
			break
		}

		if tilePos.X >= 0 && tilePos.X < s.width && tilePos.Y >= 0 && tilePos.Y < s.height {
			// Destroy walls or coins
			if s.labyrinth[tilePos.Y][tilePos.X] == 'W' || s.labyrinth[tilePos.Y][tilePos.X] == 'C' {
				s.labyrinth[tilePos.Y][tilePos.X] = 'E'
			}
		}
	}
	return stream.SendAndClose(&pb.Empty{})
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	labServer := NewLabyrinthServer()

	pb.RegisterLabyrinthServer(s, labServer)

	fmt.Println("Labyrinth server running on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
