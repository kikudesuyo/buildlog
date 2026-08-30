INSERT INTO apps (id, slug, name, category, tags, description, icon, icon_url, demo_url, code_url)
VALUES (
    6,
    'arumonogohan',
    'あるものごはん',
    'AI / Recipe',
    '["Go", "GCP Cloud Run", "Google Gemini API", "LINE Messaging API"]',
    '食べたい料理のカテゴリと使用したい食材をLINEで送信すると、レシピを提案してくれるサービスです。',
    'restaurant',
    '',
    'https://lin.ee/pLzG7zn',
    'https://github.com/kikudesuyo/arumonogohan-app'
)
ON CONFLICT (id) DO NOTHING;
