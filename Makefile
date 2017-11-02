
test:
	-dropdb pgschema
	createdb pgschema -E UTF8
	go test ./...

.PHONY: test
