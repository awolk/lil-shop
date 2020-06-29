#!/bin/bash
set -euo pipefail

DEVICE_NAME="lil-shop-dev-docker"
echo "Fetching webhook secret"

ss_file="$(mktemp)"
stripe listen --api-key $STRIPE_SECRET_KEY --device-name $DEVICE_NAME \
  -f http://localhost:3000/webhook \
  2>"${ss_file}" 1>/dev/null  &
while ! < "${ss_file}" grep -o "whsec_[^ ]*" >/dev/null; do sleep 1; done
export STRIPE_ENDPOINT_SECRET=$(< "${ss_file}" grep -o "whsec_[^ ]*")
rm "${ss_file}"

echo "Fetched webhook secret, starting server"
reflex -r "(\.go$|go\.mod)" -s go run .