CREATE TABLE "todos" (
    "id" bigserial NOT NULL,
    "task" character varying NOT NULL,
    "done" boolean NOT NULL DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY ("id"),
    CONSTRAINT "todos_task_key" UNIQUE ("task")
);

INSERT INTO todos (task) VALUES ('test-todo1');
INSERT INTO todos (task) VALUES ('test-todo2');
INSERT INTO todos (task) VALUES ('test-todo3');