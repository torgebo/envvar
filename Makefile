.PHONY: test
test:
	GOEXPERIMENT=loopvar go vet ./...
	GOEXPERIMENT=loopvar go test ./...

.PHONY: fmt
fmt:
	@go install honnef.co/go/tools/cmd/staticcheck@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	gofmt -s -w .
	goimports -w .
	staticcheck ./...



