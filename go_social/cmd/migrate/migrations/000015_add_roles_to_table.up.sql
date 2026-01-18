CREATE TABLE IF NOT EXISTS roles (
    id BIGSERIAL PRIMARY KEY, 
    name VARCHAR(255) NOT NULL UNIQUE, 
    level int NOT NULL DEFAULT 0,
    description TEXT
);

INSERT INTO roles (name, description, level) VALUES
    ('user', 'A user can create posts and comments', 1),
    ('moderator', 'A moderator can update posts and comments', 2),
    ('admin', 'An admin can create, update or delete posts and comments', 3);
