#! /bin/bash

set -e

PROTO_FILES=$(find . -type f -name "*.proto" | xargs)

for proto_file in $PROTO_FILES ; do
  protoc --go_out=go --go_opt=paths=source_relative --go-grpc_out=go --go-grpc_opt=paths=source_relative $proto_file
  echo "Generated code for $proto_file"
done