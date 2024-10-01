# Paths
PROTO_DIR=./proto
GO_OUT_DIR=./hypothesis-interface

# Tools
PROTOC=protoc
PROTOC_GEN_GO=$(shell which protoc-gen-go)

# Files
PROTO_FILES=$(PROTO_DIR)/*.proto

# Default target
all: go_generate

# Generate Go code from protobuf files
go_generate:
	mkdir -p $(GO_OUT_DIR)
	$(PROTOC) --go_out=$(GO_OUT_DIR) --go-grpc_out=$(GO_OUT_DIR) \
		--proto_path=$(PROTO_DIR) $(PROTO_FILES)

# Clean generated files
clean:
	rm -rf $(GO_OUT_DIR)

# Install protoc-gen-go (if needed)
install:
	go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

go_install:
	go mod tidy
	go mod vendor

.PHONY: all generate clean install
