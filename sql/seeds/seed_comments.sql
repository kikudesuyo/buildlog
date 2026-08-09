WITH seed(title, content) AS (
    VALUES
        ('Webエンジニア就職を実現させるために必要なこと', '自分で手を動かして学ぶことの大切さが伝わる内容でした。'),
        ('Webエンジニア就職を実現させるために必要なこと', '小さなアプリから継続して作ってみようと思います。'),
        ('静かなるシステムの構築：Rustによるメモリ安全性の再考', '安全性と開発体験のバランスについて考えさせられました。')
), target_posts AS (
    SELECT DISTINCT ON (posts.title) posts.id, posts.title
    FROM posts
    JOIN seed ON seed.title = posts.title
    WHERE posts.deleted_at IS NULL
    ORDER BY posts.title, posts.id
)
INSERT INTO comments (post_id, content)
SELECT target_posts.id, seed.content
FROM seed
JOIN target_posts ON target_posts.title = seed.title
WHERE NOT EXISTS (
    SELECT 1
    FROM comments
    WHERE comments.post_id = target_posts.id
      AND comments.content = seed.content
);
