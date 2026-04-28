#!/usr/bin/env bash
set -euo pipefail

# Requires golangci-lint: https://golangci-lint.run/welcome/install/

echo "=== go vet ==="
go vet ./...

echo "=== go test ==="
go test ./...

echo "=== golangci-lint ==="
golangci-lint run ./...

echo ""
echo "All checks passed."
