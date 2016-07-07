#!/bin/bash
DNS="$1"
SERVICE_JSON=mysql-1.0.0.json
INSTANCE_JSON=mysql-1.0.0.json
NEW_SERVICE_JSON=mysql-1.0.1.json
NEW_INSTANCE_JSON=mysql-1.0..json
INSTANCE_ID=mysql-instance-1
curl -H "Content-Type: application/json" -XPOST -d @//home/ubuntu/src/github.com/hpcloud/hdp-resource-manager/bootstrap/dev/examples/services/$SERVICE_JSON $DNS:8111/v1/services
if [ $? -eq 0 ]; then 
  echo "Create service submitted"
  curl -H "Content-Type: application/json" -XPOST -d @//home/ubuntu/src/github.com/hpcloud/hdp-resource-manager/bootstrap/dev/examples/instances/$INSTANCE_JSON $DNS:8111/v1/instances
  if [ $? -eq 0 ]; then 
    echo "Create instance submitted"
    curl -H "Content-Type: application/json" -XPOST -d @//home/ubuntu/src/github.com/hpcloud/hdp-resource-manager/bootstrap/dev/examples/services/$NEW_SERVICE_JSON $DNS:8111/v1/services
    if [ $? -eq 0 ]; then 
      echo "New Service definition submitted"
      #curl -H "Content-Type: application/json" -XPATCH -d @//home/ubuntu/src/github.com/hpcloud/hdp-resource-manager/bootstrap/dev/examples/instances/$NEW_INSTANCE_JSON $DNS:8111/v1/instances/$INSTANCE_ID
      #if [ $? -eq 0 ]; then 
	#echo "Patch Request submitted"
      #fi 
    fi 
  fi
fi 
