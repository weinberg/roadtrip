#!/bin/bash

PROJECT_BASE=`git rev-parse --show-toplevel`
cd $PROJECT_BASE/docker/playerServer/cert

rm *.pem

openssl req -x509 -sha256 -nodes -days 365 -newkey rsa:2048 -keyout key.pem -out cert.pem -subj "/CN=playerserver-gprc.insofar.com/O=Player Server GPRC"
