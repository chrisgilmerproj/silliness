# Numista Golang Client

## Getting started

Download the Numista swagger docs from: https://en.numista.com/api/doc/index.php

Ensure you have generated an API key and set the env var `NUMISTA_API_KEY`.

```sh
brew install swagger-codegen
swagger-codegen generate -i swagger.yaml -l go -o ./src/swagger
```

Run the code:

```sh
go run ./...
```
