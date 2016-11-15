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
    column "views" {
      type = "text"
    }
    index "users_email_key" {
      on = [ "lower(email)" ]
      unique = true
    }
  }
}
