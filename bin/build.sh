#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd $PROJECT_BASE/go/roadTrip

docker build -f docker/playerServer/Dockerfile . -t roadtrip/player-server
docker build -f docker/mapServer/Dockerfile . -t roadtrip/map-server
