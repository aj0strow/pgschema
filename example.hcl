
schema "v1" {
  table "users" {
    column "email" {
      type = "text"
    }
    column "name" {
      type = "citext"
    }
  }
}
