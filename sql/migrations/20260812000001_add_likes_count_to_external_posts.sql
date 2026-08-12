ALTER TABLE external_posts
    ADD COLUMN likes_count BIGINT NOT NULL DEFAULT 0;
