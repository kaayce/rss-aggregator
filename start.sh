#!/bin/bash

# Perform 'go mod tidy' to clean up the go.mod and go.sum files
go mod tidy

# Update the vendor directory with all dependencies
go mod vendor

# Build the project
go build

#  user permissions
chmod +x rss-aggregator

# execute the resulting binary
./rss-aggregator