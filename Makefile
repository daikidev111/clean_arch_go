## local install
.PHONY: local-install
local-install:
	go install golang.org/x/tools/cmd/goimports
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.42.1

## fmt
.PHONY: fmt
fmt:
	goimports -w -local "22dojo-online" cmd/ pkg/
	gofmt -s -w cmd/ pkg/

## lint
.PHONY: lint
lint:
	golangci-lint run -v cmd/... pkg/...
