CREATE TABLE goal_periods (
    id BIGSERIAL PRIMARY KEY,
    period_type TEXT NOT NULL,
    starts_at DATE NOT NULL,
    ends_at DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_goal_period_type_start UNIQUE (period_type, starts_at)
);

CREATE TABLE goals (
    id BIGSERIAL PRIMARY KEY,
    period_id BIGINT NOT NULL,
    title TEXT NOT NULL,
    target_value INTEGER NOT NULL CHECK (target_value > 0),
    progress_value INTEGER NOT NULL DEFAULT 0 CHECK (progress_value >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_goals_period FOREIGN KEY (period_id) REFERENCES goal_periods(id) ON DELETE CASCADE,
    CONSTRAINT chk_goals_progress_not_over_target CHECK (progress_value <= target_value)
);

CREATE INDEX idx_goals_period_id ON goals(period_id);
