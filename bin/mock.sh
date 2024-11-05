#!/usr/bin/env bash

set -eo pipefail

# Install this version of mockgen
# go install go.uber.org/mock/mockgen@latest

# Define the required mockgen version
REQUIRED_MOCKGEN_VERSION="v0.4.0"

# Function to check if a version string is greater than or equal to a required version
version_ge() {
  local required="$1"
  local installed="$2"

  # Use awk to compare version numbers
  awk -v installed="$installed" -v required="$required" 'BEGIN { split(installed, a, "."); split(required, b, "."); for (i=1; i<=3; i++) { if (a[i] < b[i]) { exit 1 } if (a[i] > b[i]) { exit 0 } } exit 0 }'
}

# Check mockgen version
installed_mockgen_version=$(mockgen -version || echo "mockgen not found")
if [[ "$installed_mockgen_version" == "mockgen not found" ]]; then
  echo "Error: mockgen not found. Please install mockgen (go install go.uber.org/mock/mockgen@$REQUIRED_MOCKGEN_VERSION)"
  exit 1
fi

# Check if the installed mockgen version is at least the required version
if ! version_ge "$REQUIRED_MOCKGEN_VERSION" "$installed_mockgen_version"; then
  echo "Error: mockgen version $installed_mockgen_version is not compatible. Please install mockgen $REQUIRED_MOCKGEN_VERSION (go install go.uber.org/mock/mockgen@$REQUIRED_MOCKGEN_VERSION) or higher."
  exit 1
fi

PKG_NAME=github.com/mughieams/evermos-assessment

generate_mock() {
  if `grep -q '^type.*interface {$' ${file}`; then
    dest=${file//app\//}
    [ ! -d $(dirname test/mock/${dest}) ] && mkdir -p $(dirname test/mock/${dest})
    if [[ $1 == false ]]; then
      mockgen -source=${file} -destination=test/mock/${dest}
    else
      mockgen -source=${file} | goimports -local ${PKG_NAME} > test/mock/${dest}
    fi
  fi
}

for file in `find app/common -name '*.go' | grep -v proto | grep -v /internal/`; do
  generate_mock $1
done

for file in `find app/delivery -name '*.go' | grep -v proto | grep -v /internal/`; do
  generate_mock $1
done

for file in `find app/usecase -name '*.go' | grep -v proto | grep -v /internal/`; do
  generate_mock $1
done

for file in `find app/repository -name '*.go' | grep -v proto | grep -v /internal/`; do
  generate_mock $1
done
