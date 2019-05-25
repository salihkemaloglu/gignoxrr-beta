#!/bin/bash
echo "Deleting empty folders..."
docker rmi $(docker images -f "dangling=true" -q) -f
echo "Git reset commits..."
git hard --reset
echo "Git pull current repository..."
git pull
echo "Building api project..."
docker build -t gignox_rr .
echo "Run project..."
docker run --restart=always -d -p 8901:8901 gignox_rr
