# `pgschema`

Schema as code. Write your desired schema, check it into source control, and let `pgschema` take it from there. 

### How It Works

1. Load new "A" database spec from configurable source. It can be a spec file. 
2. Load old "B" database spec from a source, usually the postgres information schema.
3. Match each element of the two database trees into a combined match tree. 
4. Walk the match tree to determine necessary changes. 
5. Filter and reorder changes depending on settings. 
6. Execute changes one by one, either printing if in dry run mode, or executing otherwise. 

### PSQL Source

`pgschema` can load your existing database spec using the postgresql information tables. It can only load the supported features. Only a subset of postgresql features work with `pgschema`.

```
information_schema.schemata
information_schema.tables
information_schema.columns
```

### HCL Source

Load your database spec from a Hashicorp Configuration Language file. Example syntax below.

```hcl
schema "public" {
    table "customers" {
        column "id" {
           type = "text"
        }
        column "email" {
           type = "text"
        }
    }
    table "products" {
        column "sku" {
           type = "text"
        }
    }
}
```
