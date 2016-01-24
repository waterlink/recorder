set -e
set -x

MAX_WAIT_RETRIES=20
RECORDER_ENDPOINT=${RECORDER_ENDPOINT:-"http://localhost:9977"}

wait_for_port() {
  retries=0
  while ! nc -z localhost $1 && [[ $retries -lt $MAX_WAIT_RETRIES ]]; do
    sleep 0.1
    retries=$((retries+1))
  done
  [[ $retries -lt $MAX_WAIT_RETRIES ]]
}

setup() {
  [[ -z "$NO_GOBUILD" ]] && go build
  ./recorder reset
  ./recorder daemon &
  wait_for_port 9977
}

cleanup() {
  curl -X_TERMINATE localhost:9977 || true
  ./recorder reset
}

TEST() {
	:
	: "		=== TEST: $@ ==="
	:
}

GET() {
	path=$1
	shift
	curl $RECORDER_ENDPOINT$path "$@"
}

POST() {
	path=$1
	shift
	curl -XPOST $RECORDER_ENDPOINT$path "$@"
}

setup
trap cleanup EXIT
