#!/bin/bash

kubectl apply -f k8s/mapSeedJob/deployment.yaml
# once it's started you have to do
# kubectl replace --force -f k8s/mapSeedJob/deployment.yaml