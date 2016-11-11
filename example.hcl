extension "uuid-ossp" {}

schema "v1" {
  table "users" {
    column "id" {
      type = "text"
      primary_key = true
    }
    column "email" {
      type = "text"
    }
    column "name" {
      type = "text"
      not_null = true
    }
    column "view_count" {
      type = "integer"
      cast_type_using = "views::integer"
    }
  }
}
