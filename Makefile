export GO111MODULE=on

VERSION := $(shell git describe --tags --abbrev=0 2>/dev/null)
COMMIT  := $(shell git rev-parse --short HEAD)
DATE    := $(shell date "+%Y-%m-%d %H:%M:%S")
LDFLAGS := -X "main.version=${VERSION}"
LDFLAGS += -X "main.commit=${COMMIT}"
LDFLAGS += -X "main.date=${DATE}"

.PHONY: all
all: deps qwy

qwy: $(shell find . -name "*.go") pkg/cmd/init.zsh
	go build -v --ldflags='${LDFLAGS}'

.PHONY: deps
deps:
	go mod tidy

.PHONY: test
test: deps
	go test ./...

.PHONY: lint
lint:
	golangci-lint run -v ./...

.PHONY: clean
clean:
	${RM} qwy
