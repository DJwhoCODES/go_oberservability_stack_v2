PROTOC_GEN_GO := $(shell which protoc-gen-go)
PROTOC_GEN_GO_GRPC := $(shell which protoc-gen-go-grpc)

proto:
	@echo "Generating all protos..."
	@./scripts/gen-proto.sh