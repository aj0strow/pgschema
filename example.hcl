extension "uuid-ossp" {}

schema "v1" {
  table "users" {
    column "email" {
      type = "text"
      primary_key = true
    }
    column "name" {
      type = "text"
      not_null = true
    }
    column "views" {
      type = "integer"
      cast_type_using = "views::integer"
    }
  }
}
