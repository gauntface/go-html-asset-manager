.PHONY: build clean gomodgen format

build: clean format
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/htmlassets ./cmds/htmlassets/htmlassets.go
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/genimgs ./cmds/genimgs/genimgs.go

# NOTE: Add the `-test.v` flag for verbose logging.
test: gomodget build
	mkdir -p coverage/
	go test ./... -covermode=atomic -coverprofile ./coverage/cover.out
	go tool cover -html=./coverage/cover.out -o ./coverage/cover.html

clean:
	rm -rf ./bin ./vendor Gopkg.lock

gomodget:
	go get -v all

format:
	go fmt ./...
