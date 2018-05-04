#!/usr/bin/env sh

GOOS=darwin GOARCH=386 go build -o binaries/ressh.app
GOOS=windows GOARCH=386 go build -o binaries/ressh.exe
GOOS=linux GOARCH=386 go build -o binaries/ressh