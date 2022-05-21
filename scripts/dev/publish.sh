#!/bin/sh

version=$(go run . raw-version)

GOPROXY=proxy.golang.org go list -m "github.com/talwat/indiepkg@$version"