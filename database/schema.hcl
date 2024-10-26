schema "public" {
  comment = "standard public schema"
}

table "todos" {
  schema = schema.public
  column "id" {
    type = serial
    null = false
  }
  column "task" {
    type = varchar(30)
    null = false
  }
  column "done" {
    type    = boolean
    default = false
    null    = false
  }
  column "created_at" {
    type    = timestamp
    null    = false
    default = sql("CURRENT_TIMESTAMP")
  }

  primary_key {
    columns = [column.id]
  }
  unique "todos_task_key" {
    columns = [column.task]
  }
}
