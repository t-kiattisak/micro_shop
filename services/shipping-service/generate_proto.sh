#!/bin/bash

if [ -z "$1" ]; then
  echo "❌ Usage: $0 <proto_filename_without_extension>"
  exit 1
fi

PROTO_NAME="$1"
PROTO_DIR="./proto"
OUT_DIR="./proto"  

if [ ! -f "$PROTO_DIR/$PROTO_NAME.proto" ]; then
  echo "❌ Error: File $PROTO_DIR/$PROTO_NAME.proto not found!"
  exit 1
fi

protoc --proto_path=$PROTO_DIR \
       --go_out=$OUT_DIR --go_opt=paths=source_relative \
       --go-grpc_out=$OUT_DIR --go-grpc_opt=paths=source_relative \
       $PROTO_DIR/$PROTO_NAME.proto

echo "✅ gRPC code for $PROTO_NAME.proto generated successfully!"