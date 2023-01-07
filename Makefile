# Makefile

PROGRAM     := influxdb-admin
VERSION     := 1.2.4
LDFLAGS     ?= "-s -w -X main.version=$(VERSION) -X main.commit=$(shell git rev-parse --short HEAD) -X 'main.build=$(shell date '+%Y-%m-%d %H:%M:%S')'"
GOBUILD_ENV = GO111MODULE=on CGO_ENABLED=0
GOBUILD     = go build -o bin/$(PROGRAM) -a -ldflags $(LDFLAGS)
GOX         = go run github.com/mitchellh/gox
TARGETS     := darwin/amd64 darwin/arm64 linux/amd64 linux/arm64 windows/amd64
DIST_DIRS   := find * -maxdepth 0 -type d -exec

.PHONY: build linux cross-build release test run statik lint down tidy clean

all: build

build:
	$(GOBUILD_ENV) $(GOBUILD)

linux:
	GOOS=linux GOARCH=amd64 $(GOBUILD_ENV) $(GOBUILD)

cross-build: clean
	$(GOBUILD_ENV) $(GOX) -ldflags $(LDFLAGS) -parallel=5 -output="bin/$(PROGRAM)-$(VERSION)-{{.OS}}-{{.Arch}}/$(PROGRAM)" -osarch='$(TARGETS)' .

release: cross-build
	( \
		cd bin && \
		$(DIST_DIRS) cp ../LICENSE {} \; && \
		$(DIST_DIRS) cp ../README.md {} \; && \
		$(DIST_DIRS) tar -zcf {}.tar.gz {} \; && \
		$(DIST_DIRS) zip -r {}.zip {} \; && \
		$(DIST_DIRS) rm -rf {} \; && \
		sha256sum * > sha256sums.txt \
	)

test:
	go test ./...

run:
	go run main.go

statik:
	go generate ./admin

lint:
	golangci-lint run --enable=golint --disable=errcheck --disable=typecheck && goimports -l -w . && go fmt ./... && go vet ./...

down:
	go list ./... && go mod verify

tidy:
	rm -f go.sum && go mod tidy -v

clean:
	rm -rf bin
