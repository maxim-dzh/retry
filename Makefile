check:
	go vet ./...
	go fmt ./...
	golint -set_exit_status $(go list ./...)
	golangci-lint run -E gofmt -E golint -E vet 
	golangci-lint run --enable maligned
	golangci-lint run -v
