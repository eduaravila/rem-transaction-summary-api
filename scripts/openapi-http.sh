#!/bin/bash

set -e

readonly service="$1"
readonly package="$2"
readonly out_dir="$3"

oapi-codegen -package "$package" -o "$out_dir/openapi.gen.go" -generate chi-server "api/openapi/$service.yml"

oapi-codegen -package "$package" -o "$out_dir/openapi_types.gen.go" -generate types "api/openapi/$service.yml"