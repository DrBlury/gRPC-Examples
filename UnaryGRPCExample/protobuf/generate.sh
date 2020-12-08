#!/bin/bash

mkdir ../generated

# gRPC Service
protoc --proto_path=proto-files --go_out=plugins=grpc:../generated ./proto-files/gRPC/*.proto