#! /usr/bin/env bash

# Generate a certificate authority

CA_KEY=temp-ca.key
CA_PEM=temp-ca.pem
CA_SRL=temp-ca.srl

# Generate the CA private key
openssl genrsa -out "${CA_KEY}" 2048

# Generate the CA root certificate
# Default subject fields
C="US"
ST="CA"
L="San Francisco"
CN="localhost"
openssl req -new -key "${CA_KEY}" -x509 -days 3652 -out "${CA_PEM}" -subj "/C=$C/ST=$ST/L=$L/O=$O/OU=$OU/CN=$CN"

# Generate devlocal cert
DEVLOCAL_CER=temp-devlocal.cer
DEVLOCAL_KEY=temp-devlocal.key
DEVLOCAL_CSR=temp-devlocal.csr

openssl req -nodes -new -keyout "${DEVLOCAL_KEY}" -out "${DEVLOCAL_CSR}" -subj "/C=$C/ST=$ST/L=$L/O=$O/OU=$OU/CN=$CN"
openssl x509 -req -in "${DEVLOCAL_CSR}" -CA "${CA_PEM}" -CAkey "${CA_KEY}" -CAcreateserial -out "${DEVLOCAL_CER}" -days 3652 -sha256
echo -n "SHA256 digest: "
openssl x509 -outform der -in "${DEVLOCAL_CER}" | openssl dgst -sha256

# Cleanup
rm -f "${DEVLOCAL_CSR}"
rm -f "${CA_SRL}"
