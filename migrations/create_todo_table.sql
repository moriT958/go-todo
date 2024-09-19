create table todos (
    id integer PRIMARY KEY,
    task varchar(100) NOT NULL UNIQUE,
    done boolean NOT NULL DEFAULT false,
    created_at timestamp NOT NULL,
    completed_at timestamp
);