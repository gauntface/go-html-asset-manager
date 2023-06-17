.PHONY: build clean gomodgen format

typescript:
	npx esbuild static/js/components/n-ham-c-lite-yt-async.ts --minify --bundle --outfile=embedassets/assets/js/components/n-ham-c-lite-yt-async.js --format=cjs
	npx esbuild static/js/components/n-ham-c-lite-vi-async.ts --minify --bundle --outfile=embedassets/assets/js/components/n-ham-c-lite-vi-async.js --format=cjs
	npx esbuild static/js/bootstrap/always-async.ts --minify --bundle --outfile=embedassets/assets/js/bootstrap/always-async.js --format=cjs

build: clean format typescript
	export GO111MODULE=on
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/htmlassets ./cmds/htmlassets/htmlassets.go
	env GOOS=linux go build -ldflags="-s -w" -o ./bin/genimgs ./cmds/genimgs/genimgs.go

# NOTE: Add the `-test.v` flag for verbose logging.
test: gomodget build
	mkdir -p coverage/
	go test ./... -covermode=atomic -coverprofile ./coverage/coverage.out
	go tool cover -html=./coverage/coverage.out -o ./coverage/coverage.html

clean:
	rm -rf ./bin ./vendor Gopkg.lock ./embedassets/assets

gomodget:
	go get -v all

format:
	go fmt ./...
