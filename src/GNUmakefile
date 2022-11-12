PLUGIN_BINARY=serve
OUTPUT_PATH=../build
VERSION=0.9.0

export GO111MODULE=on

default: build

.PHONY: clean

build:
	@env GOOS=linux GOARCH=amd64  go build -ldflags="-s -w" -o ${OUTPUT_PATH}/${PLUGIN_BINARY}_${VERSION}_linux_amd64
	#@env GOOS=windows GOARCH=amd64  go build -ldflags="-s -w" -o ${OUTPUT_PATH}/${PLUGIN_BINARY}_${VERSION}_windows_amd64
	@env GOOS=darwin GOARCH=amd64  go build -ldflags="-s -w" -o ${OUTPUT_PATH}/${PLUGIN_BINARY}_${VERSION}_darwin_amd64

clean: ## Remove build artifacts
	@rm -rf ${OUTPUT_PATH}/${PLUGIN_BINARY}*



