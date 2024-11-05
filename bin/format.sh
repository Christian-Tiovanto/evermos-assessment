#!/bin/bash

set -euo pipefail

PKG_NAME=github.com/mughieams/evermos-assessment

for f in `find . -name '*.go'`; do
  # Defensive, just in case.
  if [[ -f ${f} ]]; then
    awk '/^import \($/,/^\)$/{if($0=="")next}{print}' ${f} > /tmp/file
    mv /tmp/file ${f}
  fi
done

goimports -w -local ${PKG_NAME} $(go list -f {{.Dir}} ./...)
gofmt -s -w .
