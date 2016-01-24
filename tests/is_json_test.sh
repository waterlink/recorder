#!/usr/bin/env bash

source ./tests/lib.sh

TEST "Simple is -j statement"
POST /hello/world -d '{"hello": "world", "bga": "test"}'
./recorder is -j hello POST /hello/world = "world"
./recorder is -j bga POST /hello/world = "test"
! ./recorder is -j hello POST /hello/world = "junk content"

TEST "Simple is -j statement when no such key"
! ./recorder is -j missingKey POST /hello/world = "junk content"
! ./recorder is -j hello/stuff POST /hello/world = "junk content"

TEST "Is -j statement on Array"
POST /these/fun/numbers -d '[1, 2, 3, 5, 19, "hello world", {"hello": "world"}]'
./recorder is -j 0 POST /these/fun/numbers = 1
! ./recorder is -j 0 POST /these/fun/numbers = 75
! ./recorder is -j 0 POST /these/fun/numbers = "junk"
./recorder is -j 1 POST /these/fun/numbers = 2
./recorder is -j 2 POST /these/fun/numbers = 3
./recorder is -j 3 POST /these/fun/numbers = 5
./recorder is -j 4 POST /these/fun/numbers = 19
./recorder is -j 5 POST /these/fun/numbers = "hello world"
! ./recorder is -j 5 POST /these/fun/numbers = 42
! ./recorder is -j 6 POST /these/fun/numbers = "whatever"
./recorder is -j 6/hello POST /these/fun/numbers = "world"

TEST "Is -j statement on complex data structure"
POST /these/fun/numbers -d '{"numbers": [1, 2, [5, [8, {"what": "is that?", "and that": ["hello world", {"not mentioning": "this"}]}], 13]]}'
./recorder is -j numbers/0 POST /these/fun/numbers = 1
! ./recorder is -j numbers/0 POST /these/fun/numbers = 42
! ./recorder is -j numbers/0 POST /these/fun/numbers = "more junk"
./recorder is -j numbers/2/0 POST /these/fun/numbers = 5
! ./recorder is -j numbers/2/0 POST /these/fun/numbers = 77
./recorder is -j numbers/2/1/0 POST /these/fun/numbers = 8
./recorder is -j numbers/2/1/1/what POST /these/fun/numbers = "is that?"
! ./recorder is -j numbers/2/1/1/what POST /these/fun/numbers = 54
./recorder is -j "numbers/2/1/1/and that/0" POST /these/fun/numbers = "hello world"
! ./recorder is -j "numbers/2/1/1/and that/0" POST /these/fun/numbers = "whatever"
./recorder is -j "numbers/2/1/1/and that/1/not mentioning" POST /these/fun/numbers = "this"
./recorder is -j "numbers/2/1/1/and that/1/not mentioning" POST /these/fun/numbers = "this"
./recorder is -j "numbers/2/2" POST /these/fun/numbers = 13

TEST "Is -j statement on junk data"
POST /these/junk/endpoints -d 'whatever I am not going to be a JSON today!'
! ./recorder is -j whatever POST /these/junk/endpoints = "whatever"
