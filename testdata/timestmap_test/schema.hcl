
schema "_schema_" {
	table "users" {
		column "created_at" {
			type = "timestamp"
		}
	}
	table "customers" {
		column "created_at" {
			type = "timestamptz"
		}
	}
}
