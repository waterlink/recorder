#!/usr/bin/env bash

set -e

[[ -z "$NO_GOBUILD" ]] && go build
export NO_GOBUILD="no-go-build"

for i in $(ls ./tests/*_test.sh); do
	$i
done

echo SUCCESS
