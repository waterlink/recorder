#!/usr/bin/env bash

source ./tests/lib.sh

TEST "Simple is statement"
GET /hello/world
./recorder is GET /hello/world = ""
! ./recorder is GET /hello/world = "something"

TEST "Simple is statement with some data"
GET /hello/world\?q\=search-statement-here
./recorder is GET /hello/world = "q=search-statement-here"
! ./recorder is GET /hello/world = "junk data here"

TEST "Is statement with missing endpoints"
! ./recorder is GET /i/am/a/missing/endpoint = "no matter what"
! ./recorder is POST /hello/world = "something else"

TEST "Is statement with indexing"
GET /hello/world\?something\=different\&here\=for-you
./recorder is -i -1 GET /hello/world = "something=different&here=for-you"
./recorder is -i 0 GET /hello/world = ""
./recorder is -i 1 GET /hello/world = "q=search-statement-here"
./recorder is -i 2 GET /hello/world = "something=different&here=for-you"

TEST "Is statement with out-of-bounds indexing"
! ./recorder is -i -2 GET /hello/world = "whatever"
! ./recorder is -i 3 GET /hello/world = "not important"
! ./recorder is -i 200 GET /hello/world = "less important"

TEST "Is statement with diferent endpoints"
GET /the/different/endpoint
GET /another/endpoint\?hello\=world
./recorder is GET /hello/world = "something=different&here=for-you"
./recorder is GET /the/different/endpoint = ""
./recorder is GET /another/endpoint = "hello=world"

TEST "Is statement with different methods"
POST /hello/world -d 'the-data-is-here'
./recorder is POST /hello/world = "the-data-is-here"
! ./recorder is POST /hello/world = "wrong data"
