#!/bin/bash

export CERT_DIR=/etc/certs/${APP_HOST}
export APP_PORT=8000
export APP_ORGANIZATION=${APP_HOST}
export APP_URI="https://${APP_HOST}:${APP_PORT}"
export APP_CERT=${CERT_DIR}/${APP_HOST}.crt
export APP_KEY=${CERT_DIR}/${APP_HOST}.key
export APP_CA=${CERT_DIR}/ca.crt
if [[ ! ${APP_HOST:-} ]]; then
    echo Required variable APP_HOST not set
    echo export APP_HOST=tls-host-name
    exit 1
fi

go run secure.go util.go

# local variables:
# mode: shell-script
# end:
