#!/usr/bin/env bash

set -e

export NO_GOBUILD="no-go-build"
go build

for i in $(ls ./tests/*_test.sh); do
	$i
done

echo SUCCESS
