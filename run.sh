#!/bin/bash

# Build the Docker image
docker build --cache-from livefir -t livefir .

# Kill the existing container if it exists
if docker ps -a | grep -q livefir; then
    docker rm -f livefir
fi


# Run the container
docker run -d -p 8080:8080 --name livefir livefir

# Tail container logs
docker logs -f livefir
