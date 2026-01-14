CREATE TABLE IF NOT EXISTS user_invitations (
    token bytea PRIMARY KEY,
    user_id bigint NOT NULL
    -- optionally add FK:
    -- , CONSTRAINT user_invitations_user_fk FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
