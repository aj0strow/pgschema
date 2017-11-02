
test:
	-dropdb pgschema
	createdb pgschema -E UTF8
	go test ./...

.PHONY: test

build:
	GOOS=linux GOARCH=amd64 go build -x -o pgschema-linux-amd64 github.com/aj0strow/pgschema
	tar -cvzf pgschema-linux-amd64.tar.gz pgschema-linux-amd64
	rm pgschema-linux-amd64

.PHONY: build
