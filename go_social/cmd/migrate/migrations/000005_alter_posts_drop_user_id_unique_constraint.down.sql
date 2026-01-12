ALTER TABLE posts
ADD CONSTRAINT posts_user_id_key UNIQUE (user_id);
