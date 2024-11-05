#!/usr/bin/env bash

set -eo pipefail

# Install this version of mockgen
# go install go.uber.org/mock/mockgen@latest

PKG_NAME=github.com/mughieams/evermos-assessment

generate_mock() {
  if `test -f ${file} && grep -q '^type.*interface {$' ${file}`; then
    dest=${file//app\//}
    [ ! -d $(dirname test/mock/${dest}) ] && mkdir -p $(dirname test/mock/${dest})
    mockgen -source=${file} | goimports -local ${PKG_NAME} > test/mock/${dest}
  fi
}

for file in `git diff main... --name-only | grep .go | grep -E 'app/' | grep -v test/mock | grep -v _test.go`; do
  generate_mock
done

for file in `git ls-files -m -o | grep .go | grep -E 'app/' | grep -v test/mock | grep -v _test.go`; do
  generate_mock
done

goimports -w -local ${PKG_NAME} app/repository/postgresql
