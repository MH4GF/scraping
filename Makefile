NAME     := scraping
REVISION := $(shell git describe --always)
LDFLAGS  := -ldflags="-X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

.PHONY: help build clean run lint fmt test


help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ { printf "\033[36m%-22s\033[0m %s\n", $$1, $$NF }' $(MAKEFILE_LIST)

build: clean ## go build
	@go build -o bin/$(NAME) $(LDFLAGS)

docker-build: ## go build on Docker
	@docker build ./ -t $(NAME)

docker-run: ## run binary on Docker
	@docker run --rm -it $(NAME)

clean: ## remove binary
	@rm -f bin/$(NAME)

run: ## go run
	@go run *.go

lint: ## golint
	@golint $(go list ./... | grep -v /vendor/)

fmt: ## go fmt
	@go fmt ./...

test: ## go test
	@go test ./... -v
