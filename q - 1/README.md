# Homework-4: gRPC

```
Names : Pratham Thakkar, Swaroop C.
Roll Nos: 2021101077, 2022....
Branch : CSE, CSE
Course : Distributed Systems, Monsoon '24
```

# **_Directory Structure_**

```
📁 Q1
├── 📁 client
│   └── 📄 client.go
├── 📁 protofiles
│   └── 📄 labyrinth_grpc.pb.go
│   └── 📄 labyrinth.pb.go
│   └── 📄 labyrinth.proto
├── 📁 server
│   └── 📄 server.go
├── 📄 go.mod
├── 📄 go.sum
├── 📄 README.md
```

# Implementation Details

## Overview
In this question we have implemented a single player labyrinth game using gRPC for client-server communication. The game allows the player to navigate through a labyrinth, collect treasures, and use spells to reach the exit and win the game.

## Additional Details
At every step the user is prompted to play the game, and the following prompt is shown to the user:

```
Available commands:
1. Labyrinth Info - Get labyrinth dimensions & it's entire configuration
2. status - Get player score, health, position & spells left
3. move [DIRECTION] - Move the player (DIRECTION: up, down, left, right)
4. revelio [X] [Y] [TileType] - Use Revelio spell at given position and tile type (E / W / C)
5. bombarda [X1] [Y1] [X2] [Y2] --- [Xn] [Yn] - Use Bombarda spell at the n given positions
6. quit - to Exit the game
```

The user can enter a valid command and will be shown responses accordingly. In case an invalid command is given, an error would be thrown accordingly as well.

## Running the Code

To run the game, follow these steps: (from the q - 1 folder)

1. Start the server:
   ```
   go run ./server/server.go
   ```

2. In a separate terminal, start the client:
   ```
   go run ./client/client.go
   ```

3. Follow the prompts in the client to play the game.
