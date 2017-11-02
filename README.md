# `pgschema`

Schema as code. Write your desired schema, check it into source control, and let `pgschema` take it from there. 

## Introduction

The purpose of `pgschema` is to automate database schema updates. `pgschema` works by parsing a source file with the desired schema structure and comparing that to the existing schema using information tables. `pgschema` then diffs the two schemas and performs the minimum number of changes necessary in correct order to make the database look like the source file. 

**BEWARE**: `pgschema` is better suited for small projects or development environments. If you remove a table or column from your source file `pgschema` will delete it from the database. If you add an index to your source file `pgschema` will create a new index which can take a long time for large tables. 

```
$ go get github.com/aj0strow/pgschema
```

## Example

Create a new database called `test`.

```
$ createdb test -E UTF8
```

Create a new source schema file.

```
// schema.hcl

schema "public" {
	table "users" {
		column "id" {
			type = "text"
			primary_key = true
		}
	}
}
```

Update test database to have the same structure as the source file.

```
$ pgschema update -source schema.hcl -dsn 'dbname=test'
```

```
CREATE TABLE public.users ()
ALTER TABLE public.users ADD COLUMN id text NOT NULL
ALTER TABLE public.users ADD PRIMARY KEY (id)
```

If you run the update command again, there are no changes to make.

```
$ pgschema update -source schema.hcl -dsn 'dbname=test'
```

```
(nothing)
```

Add a new column to the source file.

```
// schema.hcl

schema "public" {
	table "users" {
		column "id" {
			type = "text"
			primary_key = true
		}
		column "first_name" {
			type = "text"
		}
	}
}
```

Run the update command to sync your changes.

```
$ pgschema update -source schema.hcl -dsn 'dbname=test'
```

```
ALTER TABLE public.users ADD COLUMN first_name text
```

If you run the update command again, there are no more changes.

```
$ pgschema update -source schema.hcl -dsn 'dbname=test'
```

```
(nothing)
```

## Heroku

Heroku integration is possible using the [custom binaries buildpack](https://github.com/tonyta/heroku-buildpack-custom-binaries). 

```
$ heroku buildpacks:add -i 1 https://github.com/tonyta/heroku-buildpack-custom-binaries#v1.0.0
```

Create a file in the root of your project called `.custom_binaries`. Go to the releases page and select the linux amd64 build and assign the link to `pgschema` command.

```
# .custom_binaries

pgschema: https://github.com/aj0strow/pgschema/releases/download/v0.1.1/pgschema-linux-amd64.tar.gz
```
