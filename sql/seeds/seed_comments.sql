INSERT INTO comments (post_id, content)
SELECT posts.id, seed.content
FROM (
    VALUES
        ('Webエンジニア就職を実現させるために必要なこと', '自分で手を動かして学ぶことの大切さが伝わる内容でした。'),
        ('Webエンジニア就職を実現させるために必要なこと', '小さなアプリから継続して作ってみようと思います。'),
        ('静かなるシステムの構築：Rustによるメモリ安全性の再考', '安全性と開発体験のバランスについて考えさせられました。')
) AS seed(title, content)
JOIN posts ON posts.title = seed.title;
