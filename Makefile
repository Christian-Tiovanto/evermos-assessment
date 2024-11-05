BUF_VERSION = v1.45.0

.PHONY: tidy
tidy:
	GO111MODULE=on go mod tidy

.PHONY: vendor
vendor:
	GO111MODULE=on go mod vendor

.PHONY: format
format:
	bin/format.sh
	buf format --write

.PHONY: lint.cleancache
lint.cleancache:
	golangci-lint cache clean

.PHONY: lint
lint: lint.cleancache
	golangci-lint run ./...

.PHONY: pretty
pretty: tidy format lint

.PHONY: test.cleancache
test.cleancache:
	go clean -testcache

.PHONY: generate.proto
generate.proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate --template "protobuf/tools/gen-go/buf.gen.yaml"

.PHONY: generate.swagger
generate.swagger:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate --template "protobuf/tools/gen-openapiv2/buf.gen.yaml"

.PHONY: generate
generate: generate.proto generate.swagger
