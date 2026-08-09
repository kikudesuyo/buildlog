WITH current_period AS (
    INSERT INTO goal_periods (period_type, starts_at, ends_at)
    VALUES (
        'monthly',
        date_trunc('month', CURRENT_DATE)::date,
        (date_trunc('month', CURRENT_DATE) + INTERVAL '1 month - 1 day')::date
    )
    ON CONFLICT (period_type, starts_at) DO UPDATE
    SET ends_at = EXCLUDED.ends_at
    RETURNING id
), seed(title, target_value, progress_value) AS (
    VALUES
        ('技術記事を公開する', 4, 2),
        ('個人開発の改善を行う', 3, 1),
        ('新しい技術を試す', 2, 1)
)
INSERT INTO goals (period_id, title, target_value, progress_value)
SELECT current_period.id, seed.title, seed.target_value, seed.progress_value
FROM current_period
CROSS JOIN seed
WHERE NOT EXISTS (
    SELECT 1
    FROM goals
    WHERE goals.period_id = current_period.id
      AND goals.title = seed.title
      AND goals.target_value = seed.target_value
      AND goals.progress_value = seed.progress_value
);
