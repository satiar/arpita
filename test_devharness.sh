#!/bin/bash
DNS=192.168.200.3
SERVICE_JSON=k8-guestbook.json
INSTANCE_JSON=guestbook.json
WORKDIR=/Users/arpita/GoProjects/src/github.com/hpcloud/hdp-resource-manager

PORT=$(curl -Ss http://192.168.200.2:8080/api/v1/namespaces/hcp/services/ipmgr | jq '.spec.ports[0].nodePort')
echo $PORT
curl -H "Content-Type: application/json" -XPOST -d @$WORKDIR/bootstrap/dev/examples/services/$SERVICE_JSON $DNS:$PORT/v1/services
if [ $? -eq 0 ]; then 
  echo "SUCCESS: Service Definition Submitted" 
  curl -H "Content-Type: application/json" -XPOST -d @$WORKDIR/bootstrap/dev/examples/instances/$INSTANCE_JSON $DNS:$PORT/v1/instances
  if [ $? -eq 0 ]; then 
    echo "SUCCESS: Instance Creation Request Submitted"
  else
    echo "FAILURE: Instance creation"
  fi
else
  echo "FAILURE: Service definition"
fi 
