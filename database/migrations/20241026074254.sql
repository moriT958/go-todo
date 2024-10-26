-- Modify "todos" table
ALTER TABLE
    "todos"
ALTER COLUMN
    "id" TYPE integer,
ALTER COLUMN
    "task" TYPE character varying(30),
ALTER COLUMN
    "created_at" TYPE timestamp;