LOCAL_BIN:=$(CURDIR)/bin
GOOSE:=$(LOCAL_BIN)/goose
MOCKERY:=$(LOCAL_BIN)/mockery

.PHONY: .bin-deps
.bin-deps:
	GOBIN=$(LOCAL_BIN) go install github.com/pressly/goose/v3/cmd/goose@latest
	GOBIN=$(LOCAL_BIN) go install github.com/vektra/mockery/v2@v2.43.2