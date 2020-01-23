
The following environment variables must be defined or the json must be configured

```
    PORT=80
    SECURE_PORT=443
    CERT_DIR=/etc/certs/${ADDRESS}
    ORGANIZATION=${ADDRESS}
    URI="https://${ADDRESS}:${PORT}"
    CERT=${CERT_DIR}/${ADDRESS}.crt
    KEY=${CERT_DIR}/${ADDRESS}.key
    CA=${CERT_DIR}/ca.crt

```
When using the script it assumes the address will be set externally

Modified from : https://github.com/duo-labs/webauthn.io

Working as a sub component under a test repo with GO111MODULE
