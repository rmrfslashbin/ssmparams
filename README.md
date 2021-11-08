# base-go
Base layout for Go projects

## Layout
```
.
├── bin
│   ├── base-go-linux-x86_64
│   └── base-go-linux-x86_64.sha256
├── cmd
│   ├── go-base
│   │   ├── main.go
│   │   └── main_test.go
│   └── hello
│       ├── main.go
│       └── main_test.go
├── Dockerfile
├── go.mod
├── go.sum
├── internal
│   ├── internal.go
│   └── internal_test.go
├── LICENSE
├── Makefile
├── pkg
│   ├── sharable.go
│   └── sharable_test.go
└── README.md
```

## Makefile
The makefile is a work in progress. Current targets:
- default: run.
- build: build all cmds in ./cmd/...
- docker-build: builds a docker image.
- install: build all cmds in ./cmd/...
- run: runs `go run cmd/hello/main.go`.
- test: runs `go test ./...`
- tidy: runs `go mod tidy`.
- update: attemtps to update go modules.

## Dockerfile
The `Dockerfile` is based on https://hub.docker.com/_/golang/. After building an image it can be run:
```
% run -it --rm github.com/rmrfslashbin/base-go:latest hello
Hello, world.

% docker run -it --rm github.com/rmrfslashbin/base-go go-base -error
hi, I'm internal!
hi, I'm exportable!
FATA[2021-11-07T17:43:36Z]/go/src/app/cmd/go-base/main.go:39 main.main.func1() main crashed error="i'm an error"
```
