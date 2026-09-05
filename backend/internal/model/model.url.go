package model

import "time"

type URL struct {
	ID          int64      `json:"id"`
	UserID      int64      `json:"user_id"`
	ShortURL    string     `json:"short_url"`
	OriginalURL string     `json:"original_url"`
	CreatedAt   time.Time  `json:"created_at"`
	ExpiresAt   *time.Time `json:"expires_at,omitempty"`
	ClickCount  int64      `json:"clicks_count"`
}

type ClickCount struct {
	ClickCount int64 `json:"clicks_count"`
}

type ShortURLExists struct {
	ShortURL bool `json:"short_url"`
}
