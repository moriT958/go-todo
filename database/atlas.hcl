env "dev" {

  src = "file://schema.hcl"
  url = "postgres://postgres:postgres@localhost:5432/mydb?search_path=public&sslmode=disable"
  dev = "docker://postgres/15/dev?search_path=public"

  migration {
    dir = "file://migrations"
  }
}
