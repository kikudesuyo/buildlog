ALTER TABLE posts ADD COLUMN status VARCHAR(20) NOT NULL DEFAULT 'draft';
UPDATE posts SET status = 'published';
