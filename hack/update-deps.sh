#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# Ensure we have everything we need under vendor/
go mod tidy
go mod vendor


