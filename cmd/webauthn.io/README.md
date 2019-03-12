
The following environment variables must be defined or the json must be configured

```
    APP_PORT=80
    APP_SECURE_PORT=443

    CERT_DIR=/etc/certs/${APP_HOST}
    APP_ORGANIZATION=${APP_HOST}
    APP_URI="https://${APP_HOST}:${APP_PORT}"
    APP_CERT=${CERT_DIR}/${APP_HOST}.crt
    APP_KEY=${CERT_DIR}/${APP_HOST}.key
    APP_CA=${CERT_DIR}/ca.crt
```

Modified from : https://github.com/duo-labs/webauthn.io

Working as a sub component under a test repo with GO111MODULE
