INSERT INTO apps (id, slug, name, category, tags, description, icon, icon_url, demo_url, code_url)
VALUES (
    5,
    'pdf-edition-v2',
    'PDF Edition v2',
    'Tool / PDF',
    '["React", "TypeScript", "Next.js", "Tailwind CSS", "Go", "GCP", "Vercel"]',
    'PDFファイルの結合と分割を行うアプリケーション。',
    'picture_as_pdf',
    '',
    'https://pdf-edition-v2.vercel.app',
    'https://github.com/kikudesuyo/pdf-edition-v2'
)
ON CONFLICT (id) DO NOTHING;
