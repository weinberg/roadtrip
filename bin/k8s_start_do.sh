#!/bin/bash

kubectl apply -f k8s/updateService/deployment.yaml
kubectl apply -f k8s/playerServer/deployment.yaml
kubectl apply -f k8s/mapServer/deployment.yaml
