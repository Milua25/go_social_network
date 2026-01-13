ALTER TABLE comments
  DROP CONSTRAINT IF EXISTS comments_post_fk,
  DROP CONSTRAINT IF EXISTS comments_user_fk;
