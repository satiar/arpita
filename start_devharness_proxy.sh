#!/bin/bash

app_version_tag=$1
registry_ip=$(docker-machine ip default)
http_proxy=http://proxy.houston.hpecorp.net:8080
no_proxy=$registry_ip,localhost,127.0.0.1
cd $WORKDIR/bootstrap/dev
INSECURE_REGISTRY=$regitry_ip:5000 HCP_REGISTRY_LOCATION=$registry_ip:5000 HTTP_PROXY=$proxy HTTPS_PROXY=$http_proxy NO_PROXY=$no_proxy APP_VERSION_TAG=$app_version_tag ./start.sh
