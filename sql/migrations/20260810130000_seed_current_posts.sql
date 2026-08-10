WITH seed(type, title, content, views, status, created_at, updated_at) AS (
    VALUES
        ('diary', 'Webエンジニア就職を実現させるために必要なこと', '2026年に新卒入社したばかりの新米エンジニアです。ソフトウェアエンジニアとして就職するためには、自走して開発するための技術力不可欠です。生成AIの台頭により就職難易度が上がる中、個人開発で技術力を高め、実務経験を積み、諦めずに挑戦し続けることが大切だと感じました。', 0, 'published', CURRENT_TIMESTAMP - INTERVAL '3 day', CURRENT_TIMESTAMP - INTERVAL '3 day'),
        ('diary', '3ヶ月間正社員として働いた感想', '社会人になってからは、開発を事業にどう活かすかを考えることが増えました。複数のプロジェクトを抱えながら優先度を立てて仕事を進める難しさがありますが、日々学びを得ながら楽しく開発しています。', 0, 'published', CURRENT_TIMESTAMP - INTERVAL '2 day', CURRENT_TIMESTAMP - INTERVAL '2 day'),
        ('tech', '静かなるシステムの構築：Rustによるメモリ安全性の再考', '現代のソフトウェア開発において、安全性はもはやオプションではありません。所有権モデルがもたらす新しい秩序と、開発者の認知負荷を軽減するための抽象化について考察します。', 1024, 'published', CURRENT_TIMESTAMP - INTERVAL '1 day', CURRENT_TIMESTAMP - INTERVAL '1 day'),
        ('tech', 'インタフェースの沈黙：ミニマリズムUIの実装戦略', '情報を削ぎ落とすことで、ユーザーの集中力を最大化する。CSS Container Queriesを活用した、文脈に応じた適応型レイアウトの設計。', 856, 'published', CURRENT_TIMESTAMP - INTERVAL '12 hour', CURRENT_TIMESTAMP - INTERVAL '12 hour'),
        ('tech', '今週のライブラリ選定：Headless UIとアクセシビリティの追求', 'スタイリングを強制しないコンポーネントが、いかにして長期的なメンテナンス性を向上させるか。Radix UIとTailwindの組み合わせ事例を詳解。', 542, 'published', CURRENT_TIMESTAMP - INTERVAL '2 hour', CURRENT_TIMESTAMP - INTERVAL '2 hour')
)
INSERT INTO posts (type, title, content, views, status, created_at, updated_at)
SELECT seed.*
FROM seed
WHERE NOT EXISTS (
    SELECT 1 FROM posts
    WHERE posts.type = seed.type AND posts.title = seed.title AND posts.content = seed.content
      AND posts.views = seed.views AND posts.status = seed.status
);
