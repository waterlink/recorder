# Recorder

Simple application that records all incoming HTTP requests and writes them down
to a filesystem in a conventional directory structure. Intended to be used in
tests.

## Installation

### Via binary release

Go to [releases](https://github.com/waterlink/recorder/releases) page and
download binary for your platform. Linux and Mac OS X are supported.

Optionally put this binary on your path (or use absolute/relative path to run
it). All examples in [usage section](#usage) assume, that binary `recorder` is
available on the path.

### Via `go get` tool

If you have properly configured Golang development environment, you can install
`recorder` via:

```bash
go get -u github.com/waterlink/recorder
```

### Via `docker`

`recorder` is available as an automated build from
[docker hub](https://hub.docker.com/r/waterlink/recorder/).

It is recommended to use a [runner script](/dockerfiles/runner.sh) as a
`recorder` on your path. Assuming `~/bin` is on your path:

```bash
curl -L https://github.com/waterlink/recorder/raw/master/dockerfiles/runner.sh > ~/bin/recorder
```

Then all examples from [usage section](#usage) should still work without
modification.

## Usage

### Start the recorder daemon

```bash
recorder daemon
```

By default it uses port `9977` and binds to `0.0.0.0`.

This can be changed by providing `-l, --listen` option, e.g.:

```bash
recorder daemon -l 127.0.0.1:3789   # or --listen=127.0.0.1:3789
```

### Get last recorded data by method & URL

```bash
recorder get POST /api/v2/user
```

### Get specific recorded data by method, URL & index

Use `-i, --index` option for that:

```bash
recorder get -i 2 POST /api/v2/user    # or --index=2
```

### Reset recorder

*NOTE: not implemented yet*

```bash
recorder reset
```

### Terminating daemon

It is enough to send `_TERMINATE` request to a server:

```bash
curl -X_TERMINATE localhost:9977
```

### Expectations

#### Expect equal last recorded data

*NOTE: not implemented yet*

```bash
recorder is POST /api/v2/user = "email=john@example.org&password=welcome"
```

`is` supports `-i, --index` option.

When expectation succeeds, `is` command returns exit code `0`.

When expectation fails, `is` command returns exit code `1`.

#### Expect JSON path to be equal

*NOTE: not implemented yet*

```bash
recorder is --json-path=user/email POST /api/v2/user = "email=john@example.org"
```

This will fail with exit code `126`, if actual data was not a parse-able JSON.

### Usage design TODO

- Design header manipulation commands/options.

## Contributing

1. Fork it ( https://github.com/waterlink/recorder/fork )
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
4. Push to the branch (git push origin my-new-feature)
5. Create a new Pull Request

## Contributors

- [waterlink](https://github.com/waterlink) Oleksii Fedorov - creator,
  maintainer

## License

MIT license. [Read here](/LICENSE).
