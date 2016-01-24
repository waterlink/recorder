#!/usr/bin/env bash

RECORDER_PORT=${RECORDER_PORT:-9977}
RECORDER_IMAGE=${RECORDER_IMAGE:-waterlink/recorder}

need_user="-u $(id -u):$(id -g)"
if [[ "$(uname)" = "Darwin" ]]; then
  need_user=
fi

daemon_opts="-it --rm"
if [[ "$1" = "daemon" ]]; then
  daemon_opts="-d -p ${RECORDER_PORT}:9977"
fi

dir_opts="-v /tmp/.recorder:/tmp/.recorder -w /tmp"

mkdir -p /tmp/.recorder

docker run \
  $need_user \
  $dir_opts \
  $daemon_opts \
  $RECORDER_IMAGE \
  recorder "$@"
