#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd $PROJECT_BASE

aws ecr --profile joshw get-login-password --region us-east-1 | docker login --username AWS --password-stdin 477748957582.dkr.ecr.us-east-1.amazonaws.com

# player-server
echo "player-server"
image=`docker -c default images | grep "^roadtrip/player-server" | awk '{print $3}'`
docker -c default tag ${image} 477748957582.dkr.ecr.us-east-1.amazonaws.com/roadtrip/player-server
docker -c default push 477748957582.dkr.ecr.us-east-1.amazonaws.com/roadtrip/player-server

# map-server
echo "map-server"
image=`docker -c default images | grep "^roadtrip/map-server" | awk '{print $3}'`
docker -c default tag ${image} 477748957582.dkr.ecr.us-east-1.amazonaws.com/roadtrip/map-server
docker -c default push 477748957582.dkr.ecr.us-east-1.amazonaws.com/roadtrip/map-server

# update-service
echo "update-service"
image=`docker -c default images | grep "^roadtrip/update-service" | awk '{print $3}'`
docker -c default tag ${image} 477748957582.dkr.ecr.us-east-1.amazonaws.com/roadtrip/update-service
docker -c default push 477748957582.dkr.ecr.us-east-1.amazonaws.com/roadtrip/update-service
