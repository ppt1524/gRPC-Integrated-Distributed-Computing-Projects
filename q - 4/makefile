# Makefile for Python gRPC Live Document Editor

.PHONY: clean proto server logs client1 client2 web web2

# Directories
PROTO_DIR=./proto
SERVER_DIR=./server
LOGGER_DIR=./logs
CLIENT_DIR=./web


# Client IDs
CLIENT1_ID=8080
CLIENT2_ID=8081



# Log File
LOG_FILE=./document_changes.log

# Default target
all: clean proto

# Clean generated protobuf files and logs
clean:
	@echo "Cleaning up generated protobuf files and logs..."
	rm -f $(PROTO_DIR)/document_pb2.py
	rm -f $(PROTO_DIR)/document_pb2_grpc.py
	@echo "Cleaned."

# Generate protobuf Python files
proto:
	@echo "Generating protobuf Python files..."
	python -m grpc_tools.protoc -I ./proto --python_out=./proto --grpc_python_out=./proto ./proto/document.proto
	@echo "Protobuf Python files generated."

# Run the gRPC server
server:
	@echo "Starting gRPC server..."
	python $(SERVER_DIR)/main.py

# Run the logging service
logs:
	@echo "Starting logging service..."
	python $(LOGGER_DIR)/main.py

# Run client 1
client1:
	@echo "Starting Client 1..."
	python $(CLIENT_DIR)/main.py --port $(CLIENT1_ID)

# Run client 2
client2:
	@echo "Starting Client 2..."
	python $(CLIENT_DIR)/main.py --port $(CLIENT2_ID)

