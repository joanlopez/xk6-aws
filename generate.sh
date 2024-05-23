#!/bin/sh

set -e

go run s3.go -flags=gen
go run eventbridge.go -flags=gen
go run kinesis.go -flags=gen