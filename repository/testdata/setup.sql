CREATE TABLE IF NOT EXISTS todos (
    id SERIAL PRIMARY KEY,
    task VARCHAR(20) UNIQUE NOT NULL,
    done BOOLEAN NOT NULL DEFAULT false,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    completed_at TIMESTAMP DEFAULT NULL
);

INSERT INTO todos (task) VALUES ('test-todo1');
INSERT INTO todos (task) VALUES ('test-todo2');
INSERT INTO todos (task) VALUES ('test-todo3');