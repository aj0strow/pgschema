extension "hstore" {}

schema "v1" {
  table "users" {
    column "email" {
      type = "text"
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
