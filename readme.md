# Homework-4: gRPC

```
Names : Pratham Thakkar
Roll No: 2021101077
Branch : CSE
Course : Distributed Systems, Monsoon '24
```

## Q - 1 (Labyrinth)

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

## Q - 2 (Distributed KNN)

# **_Directory Structure_**

```
📁 Q2
├── 📁 client
│   └── 📄 client.go
├── 📁 KNNpb
│   └── 📄 knn_grpc.pb.go
│   └── 📄 knn.pb.go
│   └── 📄 knn.proto
├── 📁 server
│   └── 📄 main.go
│   └── 📄 knn_dataset.csv
├── 📄 go.mod
├── 📄 go.sum
├── 📄 make_data.py
├── 📄 README.md (Report is written within this file)
```

# Implementation Details

## Overview
In this question we have implemented a distributed K-Nearest Neighbors (KNN) algorithm using gRPC for client-server communication. The system allows for parallel processing of KNN queries across multiple servers, each holding a portion of the dataset.

## Components

### Proto Files
- `knn.proto`: Defines the KNN service and related message types for distributed KNN implementation.

### Server
- `server.go`: Implements the KNN server, handling client queries and performing KNN calculations on its local dataset partition.
- `knn_dataset.csv`: Contains the complete KNN dataset before partitioning. (Partitioning is handled in `server.go`)

### Client
- `client.go`: Implements the client that sends KNN queries to multiple servers and aggregates results.

### Dataset
- `make_data.py`: used to generate the `knn_dataset.csv`


## Distributed KNN Approach

1. **Data Generation**: 
   - The complete dataset is generated randomly using `make_data.py`

2. **Client Query**:
   - The client sends a KNN query (point and K value) to all available servers.

3. **Server Processing**:
   - Each server performs KNN on its local dataset partition.
   - Servers return their top K nearest neighbors to the client.

4. **Result Aggregation**:
   - The client receives results from all servers.
   - It aggregates these results and performs a final KNN selection to get the global top K neighbors.

5. **Final Output**:
   - The client displays the final K nearest neighbors.

## Running the Code

To run the distributed KNN system, follow these steps (from the root directory Q2/):

1. Running the server with different ports:
   ```
   go run ./server/server.go NUM
   ```
   - The above command is used on different terminal with different values of NUM
   - NUM: is the commnad line argument that we provide to run the server with a particular port.
      - (When NUM = 0: port => 50051 + 0)
      - (When NUM = 1: port => 50051 + 1)
      - (When NUM = 2: port => 50051 + 2) & so on.

3. In a separate terminal, start the client:
   ```
   go run ./client/client.go
   ```

4. Follow the prompts in the client to input queries and receive KNN results.


### Comparative Analysis – gRPC vs MPI

1. **Communication Model**:
   - **MPI** uses tightly coupled message-passing for communication between nodes. It is synchronous and often requires the programmer to handle low-level details of the communication (e.g., message-passing between processes).
   - **gRPC** is built on top of HTTP/2, which provides more flexibility, and abstracts the communication layer, offering features like streaming and better scalability in distributed systems. It is more loosely coupled compared to MPI, which makes it easier to use across different languages and platforms.

2. **Usability**:
   - **MPI** is typically harder to set up and requires more control over the hardware resources and network.
   - **gRPC** has better support for distributed systems, microservices, and communication over a wide area network (WAN). It is more user-friendly and easier to scale up with cloud-based solutions.

3. **Scalability**:
   - **MPI** is ideal for high-performance computing environments, where communication between nodes happens on low-latency, high-bandwidth networks. It may not scale well across geographically distributed systems.
   - **gRPC** is designed for scalability over the internet, handling large numbers of distributed services with varying loads.

### Performance Analysis 

For performance evaluation, run tests with different dataset sizes and values of `k`. You can measure the following:

- **Response time**: How long it takes to find the k nearest neighbors. (When number of partitions are fixed: 5)
![Response Time vs Dataset Size](Response_Time_vs_Dataset_Size.png)
   - Explanation: In a distributed system, the response time increases with the dataset size, as there is more data to process. However, with efficient parallelization and distribution of the dataset across servers, the impact on response time is minimized. Hence, we can confirm that the system scales well and doesn’t experience exponential slowdowns as the dataset grows.
- **Scalability**: How the system behaves as the number of servers increases.
![Time vs Number of servers](Time_vs_Number_of_servers.png)
   - Explanation:  As in a well-designed distributed system, adding more servers initially reduces the workload on each server, improving response times and efficiency. However, as you add more servers, the overhead of coordinating between them (communication, network latency) starts to grow, which is why the scaling improvement becomes less dramatic after a certain number of servers. But the trend is still towards an overall improvement as the number of server increases.
- **Efficiency**: We also observerd that using streaming gave significant performance enhancement for large datasets.

## Q - 3 (My Uber)

# **_Directory Structure_**

```
📁 Q3
├── 📁 client
│   └── 📄 driver_client.go
│   └── 📄 rider_client.go
├── 📁 protofiles
│   └── 📄 myuber_grpc.pb.go
│   └── 📄 myuber.pb.go
│   └── 📄 myuber.proto
├── 📁 server
│   └── 📄 server.go
│   └── 📄 server_test.go (Contains unit tests used for testing)
├── 📄 go.mod
├── 📄 go.sum
├── 📄 Makefile
├── 📄 README.md
├── 📄 All the Certificates.
```

## Running the Code

To run the MyUber Application, follow these steps (from the root directory Q-3/):

1. Start multiple server instances in different terminals (Which will also be used for balancing the load):
   ```
   $ go run server 50051
   ```

   ```
   $ go run server 50052
   ```

   ```
   $ go run server 50053
   ```

2. In a separate terminal, start instances of the 2 clients:

   - To start rider client, you can use the following command:
        ```
        $ go run rider_client.go
        ```

    - To start driver client, you can use the following command:
        ```
        $ go run driver_client.go
        ```

5. Follow the prompts in the clients rider and driver to interact with the MyUber system. You can test the application as per your wish. All requirements for the system have been fulfilled. From the rider client, you get an option to request a ride, get it's status (if a ride is ongoing) and exit the application. From the driver end you can accept/reject ride requests.

- Q - 4 (Google Docs)

# **_Directory Structure_**

```
📁 Q4
├── 📁 logs
│   └── 📄 main.py
├── 📁 proto
│   ├── 📄 document_grpc_pb2.py
│   ├── 📄 document_grpc_pb2_grpc.py
│   └── 📄 document.proto
├── 📁 server
│   └── 📄 main.py
├── 📁 web
│   └── 📁 static
│       ├── 📁 css
│       │   └── 📄 styles.css
│       ├── 📁 js
│       │   └── 📄 app.js
│       └── 📄 index.html
    |_main.py

├── 📄 document_changes.log
├── 📄 Makefile
├── 📄 README.md
├── 📄 requirements.txt
├── 📄 toBuild.sh
```

## **Implementation Details**

### **Overview**

The Live Document system is implemented using **gRPC** and **Python** for efficient client-server communication. The system allows multiple clients to collaboratively edit a shared document in real-time.

1. **Protocol Buffers**: The `document.proto` file defines the service and message types for the Live Document system.

2. **Server**: The server maintains the current state of the document and handles client requests for updates and modifications.

3. **Client**: Clients can connect to the server, request the current document state, and send updates to modify the document.

4. **Bi-directional Streaming**: The system uses bi-directional gRPC streaming to enable real-time updates between the server and connected clients.

5. **Logging**: A logging system is implemented to track all operations and changes made to the document.

### **Key Components**

1. **Protocol Buffers** (`proto/document.proto`):
   - Defines the gRPC service and message types for the Live Document system.

2. **Server** (`server/main.py`):
   - Maintains the current state of the document.
   - Handles client connections and requests.
   - Broadcasts updates to all connected clients.
   - Communicates with the Logging Service.

3. **Client** (`client/main.py`):
   - Connects to the server using gRPC.
   - Sends document modification requests.
   - Receives and displays real-time updates from the server.

4. **Logging Service** (`logging/main.py`):
   - Receives logging requests from the server.
   - Writes document changes to `document_changes.log`.

5. **Web Interface** (`web/main.py` and `web/static/`):
   - Provides a web-based interface for users to interact with the document.
   - Uses WebSockets for real-time communication with the server.

6. **Generated Code** (`proto/document_pb2.py`, `proto/document_pb2_grpc.py`):
   - Generated from `document.proto` using `grpcio-tools`.

7. **Additional Files**:
   - `document_changes.log`: Log file recording document updates.
   - `Makefile`: Contains commands for building and running the project.
   - `requirements.txt`: Lists Python dependencies.
   - `toBuild.sh`: Script to clean and build the project.

   ## Running the Code

To run the Live Document editor, follow these steps (from the root directory Q4/):

1. Make the build script executable and run it:
   ```
   $ chmod +x toBuild.sh
   $ ./toBuild.sh
   ```

   This script will clean up previous builds and generate the necessary proto files.


2. Start the logger:
   ```
   $ make logs
   ```

3. Start the server:
   ```
   $ make server
   ```

4. In a separate terminal, start the 2 clients (can be extended to more clients, this is just for testing purposes, I've tested on more than 2 clients and the system works smoothly in that scenario as well):
   ```
   $ make client1
   ```

   ```
   $ make client2
   ```

5. Follow the prompts in the client to interact with the Live Document system.

## Makefile Usage

The Makefile provides several commands to simplify development and testing.

Note that to run the program, you also need to run the following commands in order given below (from the root directory Q4/):

- `make clean`: Removes generated files (proto files).
- `make proto`: Generates Go code from the proto files.
- `make server`: Builds and runs the server.
- `make logs`: Builds and runs the logger.
- `make client1`: Builds and runs the client1 (port: 8080).
- `make client2`: Builds and runs the client1 (port: 8081).

To use these commands, simply run `make <command>` in the terminal (from the root directory Q4/).