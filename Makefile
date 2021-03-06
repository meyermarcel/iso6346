BIN_DIR := $(GOPATH)/bin
BUILD_DIR := build
# man-pages is also defined in goreleaser.yml
MAN_DIR := man-pages
DOCS_DIR := docs
BINARY := icm

.PHONY: all
all: test lint build

.PHONY: test
test:
	go test ./...

.PHONY: lint
lint:
	golangci-lint run --enable revive --enable goimports

.PHONY: build
build:
	export CGO_ENABLED=0; go build -o $(BUILD_DIR)/$(BINARY) ./cmd/icm


# Individual commands

.PHONY: build-docs
build-docs: build
	$(shell $(BUILD_DIR)/$(BINARY) misc man $(DOCS_DIR)/$(MAN_DIR)/man1)
	$(shell $(BUILD_DIR)/$(BINARY) misc markdown $(DOCS_DIR))

.PHONY: install
install: build
	cp $(BUILD_DIR)/$(BINARY) $(BIN_DIR)/$(BINARY)

.PHONY: clean
clean:
	go clean -x -testcache
	rm -rf $(BUILD_DIR)

.PHONY: fmt
fmt:
	goimports -w $(shell find . -type f -name '*.go')

.PHONY: update-owners
update-owners: build
	echo '{}' > $(HOME)/.icm/data/owner.json && ./$(BUILD_DIR)/$(BINARY) update && cp $(HOME)/.icm/data/owner.json data/file/owner.json