-- 既存の表示形式（例: `1,024 views`）から数値だけを取り出して BIGINT に変換する。
ALTER TABLE posts
    ALTER COLUMN views DROP DEFAULT;

ALTER TABLE posts
    ALTER COLUMN views TYPE BIGINT
    USING COALESCE(NULLIF(regexp_replace(views, '[^0-9]', '', 'g'), '')::BIGINT, 0);

ALTER TABLE posts
    ALTER COLUMN views SET DEFAULT 0,
    ALTER COLUMN views SET NOT NULL;
