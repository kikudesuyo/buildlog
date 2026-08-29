CREATE TABLE learnings (
    id BIGSERIAL PRIMARY KEY,
    period_type TEXT NOT NULL CHECK (period_type IN ('daily', 'weekly', 'monthly')),
    period_start DATE NOT NULL,
    period_end DATE NOT NULL,
    content TEXT NOT NULL,
    level TEXT CHECK (level IS NULL OR level IN ('learned', 'understood', 'applied', 'explainable')),
    generated_by TEXT NOT NULL CHECK (generated_by IN ('user', 'ai')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT chk_learnings_period_dates CHECK (period_end >= period_start),
    CONSTRAINT chk_learnings_daily_level CHECK (period_type = 'daily' OR level IS NULL),
    CONSTRAINT chk_learnings_daily_dates CHECK (period_type <> 'daily' OR period_start = period_end)
);

CREATE INDEX idx_learnings_period ON learnings(period_type, period_start, period_end);
