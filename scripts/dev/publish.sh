#!/bin/sh

version=$(go run . raw-version)

go get -u ./...
go mod tidy
GOPROXY=proxy.golang.org go list -m "github.com/talwat/indiepkg@$version"