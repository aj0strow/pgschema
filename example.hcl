extension "hstore" {}

schema "v1" {
  table "users" {
    column "email" {
      type = "text"
    }
    column "name" {
      type = "text"
    }
    column "views" {
      type = "integer"
      cast_type_using = "$name::integer"
    }
  }
}
