#!/bin/bash

# Change to the root directory of the project
cd "$(dirname "$0")/.." || exit

# Set environment variables for cross-compilation
export GOOS=windows
export GOARCH=amd64

# Build the Windows executable
go build -o builds/go-ip-recon.exe ./cmd/main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "Windows executable created: builds/go-ip-recon.exe"
else
    echo "Build failed"
fi