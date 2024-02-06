generate:
	go generate ./...

lint: generate
	gofmt -s -w .
	goimports -local github.com/mt-inside/polybar-lmsensors -w .
	go vet ./...
	staticcheck ./...
	golangci-lint run ./... # TODO: --enable-all

run *ARGS: lint
	go run . {{ARGS}}

# This isn't a "static" build, but that's ok.
# The go runtime gets statically linked anyway, so this only depends on libc, libsensors, etc
# ie it has reasonable deps, but doesn't need go to be installed on the user's system
# Static linking when using cgo is really hard, not least you need STATIC versions of the c libraries (like libsensors), which themselves have libc baked in
build: lint
	go build .

install: lint
	go install .
