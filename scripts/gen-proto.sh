#!/bin/bash
set -e

PROTO_DIR=proto

echo "Generating User Proto..."
protoc \
  --go_out=paths=source_relative:./user-service/internal/gen/user \
  --go-grpc_out=paths=source_relative:./user-service/internal/gen/user \
  $PROTO_DIR/user.proto

echo "Generating Post Proto..."
protoc \
  --go_out=paths=source_relative:./post-service/internal/gen/post \
  --go-grpc_out=paths=source_relative:./post-service/internal/gen/post \
  $PROTO_DIR/post.proto

echo "Generating Notification Proto..."
protoc \
  --go_out=paths=source_relative:./notification-service/internal/gen/notification \
  --go-grpc_out=paths=source_relative:./notification-service/internal/gen/notification \
  $PROTO_DIR/notification.proto

echo "✅ All protos generated!"