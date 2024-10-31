package main

import (
	"bufio"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"os"
	"strings"
	"time"

    pb "q-1/labyrinthpb" // Use the correct package path
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to server: %v", err)
	}
	defer conn.Close()

	client := pb.NewLabyrinthClient(conn)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Choose an action: Labyrinth Info, move, revelio, bombarda, status, quit")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)

		switch action {
		case "Labyrinth Info":
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			resp, err := client.GetLabyrinthInfo(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Error getting labyrinth info: %v", err)
			}
			fmt.Printf("Labyrinth Info - width: %d, height: %d\n", resp.Width, resp.Height)
			fmt.Printf("The entire Labyrinth is as follows:\n")
			for i := range resp.Labyrinth{
				fmt.Println(resp.Labyrinth[i])
			}
		
		case "move":
			fmt.Print("Enter direction (up/down/left/right): ")
			direction, _ := reader.ReadString('\n')
			direction = strings.TrimSpace(direction)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			resp, err := client.RegisterMove(ctx, &pb.MoveRequest{Direction: direction})
			if err != nil {
				log.Fatalf("Error while moving: %v", err)
			}
			fmt.Println("Move result:", resp.Status)
			if resp.Status == "Death & Game Over"  || resp.Status == "victory" || resp.Status == "hit wall & victory" {
				return
			}
			

		case "revelio":
			fmt.Print("Enter target x, y: ")
			var x, y int32
			fmt.Scan(&x, &y)

			fmt.Print("Enter TileType: ")
			var TyleType string
			fmt.Scan(&TyleType)

			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			stream, err := client.Revelio(ctx, &pb.RevelioRequest{X: x, Y: y, TileType: TyleType})
			if err != nil {
				log.Fatalf("Error casting revelio: %v", err)
			}

			for {
				tile, err := stream.Recv()
				if err != nil {
					break
				}
				fmt.Printf("Found tile at (%d, %d)\n", tile.X, tile.Y)
			}

		case "bombarda":
			fmt.Println("Enter tile positions to destroy (format: x y). Enter 'done' to finish.")
			stream, err := client.Bombarda(context.Background())
			if err != nil {
				log.Fatalf("Error casting bombarda: %v", err)
			}

			for {
				fmt.Print("Enter tile position (x y): ")
				var x, y int32
				_, err := fmt.Scan(&x, &y)
				if err != nil {
					break
				}

				stream.Send(&pb.TilePosition{X: x, Y: y})

				input, _ := reader.ReadString('\n')
				if strings.TrimSpace(input) == "done" {
					break
				}
			}
			stream.CloseAndRecv()

		case "status":
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()

			resp, err := client.GetPlayerStatus(ctx, &pb.Empty{})
			if err != nil {
				log.Fatalf("Error getting player status: %v", err)
			}
			fmt.Printf("Player Status - Score: %d, Health: %d, Position: (%d, %d), Spells left: %d\n", resp.Score, resp.Health, resp.X, resp.Y, resp.RemainingSpells)

		case "quit":
			fmt.Println("Exiting the game.")
			return

		default:
			fmt.Println("Invalid action. Try again.")
		}
	}
}
