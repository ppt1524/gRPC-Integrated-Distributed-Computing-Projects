#!/bin/bash

# toBuild.sh

# Exit immediately if a command exits with a non-zero status
set -e

echo "Cleaning previous builds..."
make clean

echo "Generating protobuf files..."
make proto

echo "Build and setup completed."