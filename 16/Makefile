BIN=bin/aoc_16
GOCMD=go

.PHONY=build test clean tidy
build: tidy test
	go build -o ${BIN} main.go

test: verify
	go test -v -timeout=1000000s ./...

clean:
	rm -f ${BIN}

tidy:
	go fmt ./...
	go mod tidy -v

verify:
	go mod verify
	go vet ./....
	GOFLAGS="-buildvcs=false" go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...
