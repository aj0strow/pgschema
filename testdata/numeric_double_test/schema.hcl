
schema "_schema_" {
	table "accounts" {
		column "id" {
			type = "text"
			not_null = true
		}
		column "balance" {
			type = "double precision"
			not_null = true
			default = "0"
		}
	}
}
