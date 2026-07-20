variable "database_url" {
  type    = string
  default = getenv("DATABASE_URL")
}

env "local" {
  url = var.database_url
  dev = "docker://postgres/17/dev?search_path=public"

  schema {
    src = "file://sql/schema.sql"
  }

  migration {
    dir = "file://sql/migrations"
  }
}
