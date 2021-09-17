#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd $PROJECT_BASE

docker build -f docker/playerServer/Dockerfile go/roadTrip -t roadtrip/player-server
docker build -f docker/mapServer/Dockerfile go/roadTrip -t roadtrip/map-server
docker build -f docker/updateService/Dockerfile go/roadTrip -t roadtrip/update-service
