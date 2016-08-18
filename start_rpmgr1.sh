#!/bin/bash
docker run \
  --rm \
  -p 8113:8113 \
  -e CLUSTER_ID=12345 \
  -e RPMGR_PROVIDER=local \
  -e RPMGR_DOCKER_VERSION=1.10.3-0~trusty \
  -e NATS_HOST=$(docker-machine ip default):4222 \
  -e NATS_USERNAME=hcp \
  -e NATS_PASSWORD_FILE=/etc/secrets/nats \
  -e RPMGR_DB_TYPE=inmemory \
  -e LOG_LEVEL=debug \
  -e RPMGR_KEEP_TERRAFORM=true \
  -e RPMGR_TERRAFORM_DIR=/tf \
  -e CA_CERT_FILE=/etc/secrets/hcp_ca.crt \
  -e NTP_SERVERS="ntp1,ntp2" \
  -e HCP_USE_SYSLOG=false \
  -e NATS_USE_TLS=true \
  -e CA_KEY_FILE=/etc/secrets/server.key \
  -v $HOME/arpita/secrets:/etc/secrets \
  -v $HOME/tf:/tf \
  de5806a84230
