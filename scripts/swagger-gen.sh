#!/bin/sh

set -e

rm -rf client
rm -rf models

# docker run --rm -it  --user $(id -u):$(id -g) -v $(pwd):$(pwd)  -w $(pwd) quay.io/goswagger/swagger:v0.27.0 generate client -f https://dkron.io/swagger.yaml -t $(pwd)
docker run --rm -it  --user $(id -u):$(id -g) -v $(pwd):$(pwd)  -w $(pwd) quay.io/goswagger/swagger:v0.27.0 generate client -f $(pwd)/scripts/swagger.yaml -t $(pwd)

go mod tidy