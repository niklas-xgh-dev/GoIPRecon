#!/bin/bash

# Change to the root directory of the project
cd "$(dirname "$0")/.." || exit

# Ensure we're building for macOS
unset GOOS GOARCH

# Build the macOS executable
go build -o builds/go-ip-recon ./cmd/main.go

# Check if the build was successful
if [ $? -eq 0 ]; then
    echo "macOS executable created: builds/go-ip-recon"
else
    echo "Build failed"
fi