pkg = github.com/aj0strow/pgschema

test:
	-dropdb pgschema
	createdb pgschema -E UTF8
	go test ./...

.PHONY: test

build:
	sh ./build.sh
	
.PHONY: build
