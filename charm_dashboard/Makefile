
LDFLAGS=-extldflags=-static -s -w

.PHONY: help
help:  ## Print the help documentation
	@grep -E '^[\/a-zA-Z0-9_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: build
build: ## Build the binary
	CGO_ENABLED=0 go build -installsuffix cgo -ldflags "$(LDFLAGS)" -o bin/ssmpicker src/*.go

.PHONY: build_linux_amd64
build_linux_amd64: ## Build the binary for linux AMD64 and sha256sum file
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -installsuffix cgo -ldflags "$(LDFLAGS)" -o bin/ssmpicker_linux_amd64 src/*.go

.PHONY: build_darwin_amd64
build_darwin_amd64: ## Build the binary for darwin AMD64 and sha256sum file
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -installsuffix cgo -ldflags "$(LDFLAGS)" -o bin/ssmpicker_darwin_amd64 src/*.go

.PHONY: build_darwin_arm64
build_darwin_arm64: ## Build the binary for darwin AMD64 and sha256sum file
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -installsuffix cgo -ldflags "$(LDFLAGS)" -o bin/ssmpicker_darwin_arm64 src/*.go

.PHONY: build_all
build_all: build_linux_amd64 build_darwin_amd64 build_darwin_arm64

.PHONY: run
run: ## Run the code
	go run ./...
