ALTER TABLE posts ADD COLUMN deleted_at TIMESTAMPTZ DEFAULT NULL;
CREATE INDEX idx_posts_deleted_at ON posts(deleted_at);
