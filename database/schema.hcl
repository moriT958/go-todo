# スキーマの理想状態
table "todos" {
  schema = schema.public
  column "id" {
    null = false
    type = bigserial
  }
  column "task" {
    null = false
    type = character_varying
  }
  column "done" {
    null    = false
    type    = boolean
    default = false
  }
  column "created_at" {
    null    = false
    type    = timestamptz
    default = sql("CURRENT_TIMESTAMP")
  }
  primary_key {
    columns = [column.id]
  }
  unique "todos_task_key" {
    columns = [column.task]
  }
}
schema "public" {
  comment = "standard public schema"
}
