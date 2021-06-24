#!/bin/bash

npx swagger2openapi scripts/swagger.yaml -y -o scripts/oa3.yaml

go get github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.8.1
oapi-codegen -package client -generate=types,client -o=client/dkronc.gen.go scripts/oa3.yaml