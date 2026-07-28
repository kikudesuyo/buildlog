CREATE TABLE apps (
    id BIGSERIAL PRIMARY KEY,
    slug TEXT NOT NULL UNIQUE,
    name TEXT NOT NULL,
    category TEXT NOT NULL,
    tags JSONB NOT NULL,
    description TEXT NOT NULL,
    icon TEXT NOT NULL,
    icon_url TEXT DEFAULT '',
    demo_url TEXT DEFAULT '',
    code_url TEXT DEFAULT '',
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO apps (id, slug, name, category, tags, description, icon, icon_url, demo_url, code_url)
VALUES 
(
    1,
    'whichway',
    'Whichway',
    'Tool / Decision Support',
    '["TypeScript", "Go"]',
    '公共交通機関の乗換案内アプリです。',
    'explore',
    '/whichway-icon.svg',
    'https://whichway-6862oinh7-kikudeusyos-projects.vercel.app/',
    'https://github.com/kikudesuyo/whichway'
),
(
    2,
    'mahjong-scoreboard',
    '麻雀スコアボード管理',
    'Tool / Utility',
    '["TypeScript", "Go"]',
    '対局のスコア記録・計算・成績管理をスムーズに行える麻雀専用スコアボード。',
    'score',
    '/mahjong-icon.svg',
    'https://mahjong-scoreboard-management.vercel.app/',
    ''
),
(
    3,
    'pratan',
    'Pratan',
    'Education / Language',
    '["TypeScript", "React", "Firebase"]',
    '英単語の効率的な学習と記憶定着をサポートする語学学習アプリケーション。クイズ機能を搭載しています',
    'translate',
    '/pratan-icon.svg',
    'https://pratan-714.web.app/',
    'https://github.com/kikudesuyo/pratan'
),
(
    4,
    'econom-eye',
    'economEye',
    'Finance / Tracking',
    '["TypeScript", "React", "Firebase"]',
    '商品の価格推移を視覚的に追跡・監視する商品価格追跡ツール。最初に作成したWebアプリケーションです。',
    'monitoring',
    '/economeye-icon.svg',
    'https://economeye-d5146.web.app/',
    'https://github.com/kikudesuyo/economEye'
)
ON CONFLICT (id) DO NOTHING;
