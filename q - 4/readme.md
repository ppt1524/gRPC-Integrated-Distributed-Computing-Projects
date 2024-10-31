Names : Pratham Thakkar, Swaroop C.
Roll Nos: 2021101077, 2022...
Branch : CSE, CSE
Course : Distributed Systems, Monsoon '24
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