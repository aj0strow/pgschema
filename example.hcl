extension "uuid-ossp" {}

schema "v1" {
  table "users" {
    column "id" {
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
      prior_name = "views"
    }
    
    constraint "users_pkey" {
      primary_key = [ "id", "created_at" ]
    }
    
    
    
    constraint "users_pkey" {
      primary_key = [ "id", "created_at" ]
    }
    
    primary_key = [ "id", "created_at" ]
  }
}
