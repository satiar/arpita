#!/bin/bash
docker run \
  --rm \
  -p 4222:4222 \
  -e NATS_USERNAME=hcp \
  -e NATS_PASSWORD_FILE=/etc/secrets/nats \
  -v ~/secrets:/etc/secrets \
  hcp/nats
