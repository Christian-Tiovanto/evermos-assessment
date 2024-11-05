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

.PHONY: mock
mock:
	bin/mock.sh

.PHONY: mock.partial
mock.partial:
	bin/mock-partial.sh

.PHONY: test.unit
test.unit: test.cleancache
	go test -race ./...

.PHONY: test.cover
test.cover: test.cleancache
	go test -v -race ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html
	go tool cover -func coverage.out

.PHONY: migration
migration:
	dbmate -d ./db/migrations new $(name)

.PHONY: migration.status
migration.status:
	dbmate --wait --url $(url) status

.PHONY: migrate
migrate:
	dbmate --wait --url $(url) --no-dump-schema migrate --strict -v

.PHONY: rollback
rollback:
	dbmate --wait --url $(url) --no-dump-schema rollback -v

.PHONY: generate.proto
generate.proto:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate --template "protobuf/tools/gen-go/buf.gen.yaml"

.PHONY: generate.swagger
generate.swagger:
	go run github.com/bufbuild/buf/cmd/buf@$(BUF_VERSION) generate --template "protobuf/tools/gen-openapiv2/buf.gen.yaml"

.PHONY: generate
generate: generate.proto generate.swagger

.PHONY: sqlc
sqlc:
	sqlc generate
	goimports -w -local github.com/mughieams/evermos-assessment app/repository/postgresql/db

.PHONY: run
run:
	go run cmd/server/main.go
