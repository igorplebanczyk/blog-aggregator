-- +goose up
CREATE TABLE feeds (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    name VARCHAR NOT NULL,
    url VARCHAR NOT NULL UNIQUE,
    user_id UUID NOT NULL REFERENCES users ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);

-- +Goose down
DROP TABLE feeds;