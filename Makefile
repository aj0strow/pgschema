
test: fmt
	go test ./...

fmt:
	go fmt ./...

.PHONY: test fmt
