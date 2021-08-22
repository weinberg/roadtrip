#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd $PROJECT_BASE/docker

docker-compose up
