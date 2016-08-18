#!/bin/bash
docker run \
  --rm \
  -p 4222:4222 \
  -e NATS_USERNAME=hcp \
  -e NATS_PASSWORD_FILE=/etc/secrets/nats \
  -e NATS_USE_TLS=true \
  -e NATS=/etc/secrets/nats \
  -e NATS_TLS_KEY=/etc/sercrets/nats-tls-key \
  -e NATS_TLS_CERT=/etc/secrets/nats-tls-cert \
  -e CA_CERT_FILE=/etc/secrets/hcp_ca.crt \
  -v ~/arpita/secrets:/etc/secrets \
  hcp/nats
