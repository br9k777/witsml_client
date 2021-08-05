.PHONY: setup
setup: ## Install all the build and lint dependencies
	[ -r ${GOPATH}/bin/golangci-lint ] && rm ${GOPATH}/bin/golangci-lint
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(GOPATH)/bin v1.21.0
# 	@$(MAKE) dep

# .PHONY: cover
# cover: test ## Run all the tests and opens the coverage report
# 	go tool cover -html=coverage.txt

.PHONY: fmt
fmt: ## Run go imports on all go files
	find . -name '*.go' -not -wholename './vendor/*' | while read -r file; do goimports -w "$$file"; done

.PHONY: test ## Run all the tests
test: test1
test1:
	cd pkg/elements && go test -count=1 -run TestReadLogs

.PHONY: lint
lint: ## Run all the linters
	golangci-lint run ./...

# .PHONY: lint
# lint: ## Run all the linters
# 	gometalinter --vendor --disable-all \
# 		--enable=deadcode \
# 		--enable=ineffassign \
# 		--enable=gosimple \
# 		--enable=staticcheck \
# 		--enable=gofmt \
# 		--enable=goimports \
# 		--enable=misspell \
# 		--enable=errcheck \
# 		--enable=vet \
# 		--enable=vetshadow \
# 		--deadline=10m \
# 		./...

.PHONY: build
build: ## Build a version
	GOOS="linux" GOARCH="amd64" go build -o /tmp/print_witsml_logs ./cmd/parse_logs



.PHONY: clean
clean: ## Remove temporary files
	go clean

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'


