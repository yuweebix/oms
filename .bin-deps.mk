LOCAL_BIN:=$(CURDIR)/bin

PROTOC:=$(LOCAL_BIN)/protoc

GOOSE:=$(LOCAL_BIN)/goose
MOCKERY:=$(LOCAL_BIN)/mockery
PROTOC_GEN_GO:=$(LOCAL_BIN)/protoc-gen-go
PROTOC_GEN_GO_GRPC:=$(LOCAL_BIN)/protoc-gen-go-grpc
PROTOC_GEN_GRPC_GATEWAY:=$(LOCAL_BIN)/protoc-gen-grpc-gateway
PROTOC_GEN_OPENAPI:=$(LOCAL_BIN)/protoc-gen-openapiv2
PROTOC_GEN_VALIDATE:=$(LOCAL_BIN)/protoc-gen-validate

PROTOC_URL:=https://github.com/protocolbuffers/protobuf/releases/download/v27.2/protoc-27.2-linux-x86_64.zip
PROTOC_ZIP:=$(CURDIR)/protoc-27.2-linux-x86_64.zip

.bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.43.2

.bin-protoc-deps: .protoc
	GOBIN=$(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN=$(LOCAL_BIN) go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	GOBIN=$(LOCAL_BIN) go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest
	GOBIN=$(LOCAL_BIN) go install github.com/envoyproxy/protoc-gen-validate@latest

.protoc:
	curl -Lo $(PROTOC_ZIP) $(PROTOC_URL)
	unzip -o $(PROTOC_ZIP)
	rm -fr $(PROTOC_ZIP) readme.txt include/
