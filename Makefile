GOCMD=GO111MODULE=on go

build: ## build the go app
	go build -o tmp/base32check1 ./base32check1.go

linters-install:
	@golangci-lint --version >/dev/null 2>&1 || { \
		echo "installing linting tools..."; \
		curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.41.1; \
	}

lint: linters-install
	golangci-lint run

install:
	go install ./base32check1.go

test:
	$(GOCMD) test -cover -race ./...

test-ci:
	$(GOCMD) test -cover -race ./...

bench:
	$(GOCMD) test -bench=. -benchmem ./...

.PHONY: test lint linters-install
