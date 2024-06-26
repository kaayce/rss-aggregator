#!/bin/bash

# Perform 'go mod tidy' to clean up the go.mod and go.sum files
go mod tidy

# Update the vendor directory with all dependencies
go mod vendor
