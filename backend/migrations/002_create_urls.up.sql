CREATE TABLE urls (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    short_url VARCHAR(10) NOT NULL UNIQUE,
    original_url TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    expires_at TIMESTAMPTZ,
    clicks_count BIGINT NOT NULL DEFAULT 0,

    CONSTRAINT urls_user_constraint
        FOREIGN KEY (user_id)
        REFERENCES users(id)
        ON DELETE CASCADE 
);

CREATE INDEX urls_user_created_at_index
ON urls(user_id, created_at DESC);
