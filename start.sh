#!/bin/bash

# Run the tidy script
./tidy.sh

# Build the project
go build

#  user permissions
chmod +x rss-aggregator

# execute the resulting binary
./rss-aggregator