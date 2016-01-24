set -e
set -x

RECORDER_ENDPOINT=${RECORDER_ENDPOINT:-"http://localhost:9977"}

setup() {
  [[ -z "$NO_GOBUILD" ]] && go build
  ./recorder reset
  ./recorder daemon &
}

cleanup() {
  curl -X_TERMINATE localhost:9977
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
