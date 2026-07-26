CREATE TABLE profiles (
    id BIGSERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    subtitle TEXT NOT NULL,
    title TEXT NOT NULL,
    avatar_url TEXT NOT NULL,
    quote TEXT NOT NULL,
    bio JSONB NOT NULL,
    highlights JSONB NOT NULL,
    award TEXT DEFAULT '',
    expertise JSONB NOT NULL,
    contact_email TEXT NOT NULL,
    final_quote TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO profiles (id, name, subtitle, title, avatar_url, quote, bio, highlights, award, expertise, contact_email, final_quote)
VALUES (
    1,
    'kikudesuyo',
    'Hriomu Kikuchi — Product Engineer',
    'Creative Director',
    'https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=300&q=80',
    '「余白とは、単なる空白ではなく、思考が呼吸するための空間である。」',
    '["15年間にわたり、デザインとテクノロジーの交差点で「静寂」を追求してきました。情報の過剰な現代において、真に価値のある体験とは、引き算によってのみ生まれると信じています。", "現在は Buildlog のリードデザイナーとして、執筆者と読者が深く繋がれるデジタル空間の構築に注力しています。物理的なノートのような手触り感と、デジタルの効率性を融合させた、新しい編集体験を提案しています。"]',
    '[{"title": "Buildlog Platform", "period": "2021 — Present", "description": "次世代の執筆環境をデザイン。月間100万人のアクティブユーザーを持つプラットフォームへと成長。"}, {"title": "Mono Design Studio", "period": "2018 — 2021", "description": "創設者。12の国際的なデザイン賞を受賞し、ミニマリズムの先駆者として認知される。"}]',
    'Global Design Excellence 2023',
    '["インタラクション設計", "タイポグラフィ", "UXライティング", "ブランド戦略", "ミニマリストUI", "コンテンツ戦略"]',
    'contact@buildlog.dev',
    '「物語は、余白の中で最も力強く響く。あなたの想いを聞かせてください。」'
) ON CONFLICT (id) DO NOTHING;
