#!/bin/bash
IP=localhost
TAG=test160today2
docker tag -f hcp/rpmgr:latest $IP:5000/hcp/rpmgr:$TAG
docker tag -f hcp/ipmgr:latest $IP:5000/hcp/ipmgr:$TAG
docker tag -f hcp/uaa:latest $IP:5000/hcp/uaa:$TAG
docker tag -f hcp/nats:latest $IP:5000/hcp/nats:$TAG
docker tag -f hcp/flightrecorder:latest $IP:5000/hcp/flightrecorder:$TAG
docker tag -f hcp/configuration:latest $IP:5000/hcp/configuration:$TAG
docker tag -f hcp/redis:latest $IP:5000/hcp/redis:$TAG
docker push $IP:5000/hcp/rpmgr:$TAG
docker push $IP:5000/hcp/ipmgr:$TAG
docker push $IP:5000/hcp/uaa:$TAG
docker push $IP:5000/hcp/flightrecorder:$TAG
docker push $IP:5000/hcp/nats:$TAG
docker push $IP:5000/hcp/configuration:$TAG
docker push $IP:5000/hcp/redis:$TAG
