#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd $PROJECT_BASE

docker build -f docker/playerServer/Dockerfile . -t roadtrip/player-server
docker build -f docker/mapServer/Dockerfile . -t roadtrip/map-server
docker build -f docker/mapServerSeed/Dockerfile . -t roadtrip/map-seed-job
docker build -f docker/updateService/Dockerfile . -t roadtrip/update-service
