#!/bin/bash

app_version_tag=$1
registry_ip=$(docker-machine ip default)
cd $WORKDIR/bootstrap/dev
INSECURE_REGISTRY=$regitry_ip:5000 HCP_REGISTRY_LOCATION=$registry_ip:5000 APP_VERSION_TAG=$app_version_tag ./start.sh
