#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd $PROJECT_BASE

docker login

# player-server
echo "player-server"
image=`docker -c default images | grep "^roadtrip/player-server" | awk '{print $3}'`
docker -c default tag ${image} 0x01/roadtrip-player-server
docker -c default push 0x01/roadtrip-player-server

# map-server
echo "map-server"
image=`docker -c default images | grep "^roadtrip/map-server" | awk '{print $3}'`
docker -c default tag ${image} 0x01/roadtrip-map-server
docker -c default push 0x01/roadtrip-map-server

# update-service
echo "update-service"
image=`docker -c default images | grep "^roadtrip/update-service" | awk '{print $3}'`
docker -c default tag ${image} 0x01/roadtrip-update-service
docker -c default push 0x01/roadtrip-update-service

