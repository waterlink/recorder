#!/usr/bin/env bash

source ./tests/lib.sh

TEST "Simple get"
GET /hello/world
[[ "$(./recorder get GET /hello/world)" = "" ]]

TEST "Simple get with some data"
GET /hello/world\?name\=earth\&greeting\=hi
[[ "$(./recorder get GET /hello/world)" = "name=earth&greeting=hi" ]]

TEST "Get indexing"
GET /hello/world\?some\=testing
[[ "$(./recorder get -i -1 GET /hello/world)" = "some=testing" ]]
[[ "$(./recorder get -i 0 GET /hello/world)" = "" ]]
[[ "$(./recorder get -i 1 GET /hello/world)" = "name=earth&greeting=hi" ]]
[[ "$(./recorder get -i 2 GET /hello/world)" = "some=testing" ]]

TEST "Out-of-bounds get indexing"
! ./recorder get -i -2 GET /hello/world
! ./recorder get -i 3 GET /hello/world
! ./recorder get -i 200 GET /hello/world

TEST "Get with different endpoints"
GET /i/am/different/endpoint\?hello\=world
[[ "$(./recorder get GET /i/am/different/endpoint)" = "hello=world" ]]
[[ "$(./recorder get GET /hello/world)" = "some=testing" ]]

TEST "Get different methods"
POST /i/am/different/endpoint -d 'some=data&here=and-there'
[[ "$(./recorder get GET /i/am/different/endpoint)" = "hello=world" ]]
[[ "$(./recorder get POST /i/am/different/endpoint)" = "some=data&here=and-there" ]]
