#!/bin/bash
DNS=$1
INSTANCE_ID=mysql-instance-1
NEW_INSTANCE_JSON=mysql-1.0.1.json
curl -v -H "Content-Type: application/json" -XPATCH -d @//home/ubuntu/src/github.com/hpcloud/hdp-resource-manager/bootstrap/dev/examples/instances/$NEW_INSTANCE_JSON $DNS:8111/v1/instances/$INSTANCE_ID
if [ $? eq 0 ]
  echo "SUCCESS!"
fi 
