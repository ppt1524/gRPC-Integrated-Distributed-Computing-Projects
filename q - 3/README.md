# Homework-4: gRPC

```
Names : Pratham Thakkar, Swaroop C.
Roll Nos: 2021101077, 2022....
Branch : CSE, CSE
Course : Distributed Systems, Monsoon '24
```

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