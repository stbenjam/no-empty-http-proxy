all: noemptyhttpproxy.so noemptyhttpproxy
.PHONY: lint test

clean:
	rm -f noemptyhttpproxy.so noemptyhttpproxy

test:
	go test ./...

lint:
	golangci-lint run ./...

noemptyhttpproxy:
	go build ./cmd/noemptyhttpproxy

noemptyhttpproxy.so:
	go build -buildmode=plugin ./plugin/noemptyhttpproxy.go
