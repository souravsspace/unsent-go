#!/bin/bash
set -e

# Ensure we are in the sdk root or handle paths correctly
# This script assumes it's run from the cli/go-sdk directory or we can find the schema relative to it.

SCHEMA_PATH="../../apps/docs/public/api-reference.json"
OUTPUT_PATH="./pkg/unsent/types.go"

echo "Generating Go types from ${SCHEMA_PATH}..."

# Generate types from schema
go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest -package unsent -generate types -o "$OUTPUT_PATH" "$SCHEMA_PATH"

echo "Done. Types generated at ${OUTPUT_PATH}"
