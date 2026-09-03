package repository

import (
	"context"
	"goBitly/internal/model"
	"github.com/jackc/pgx/v5/pgxpool"
)

type URLRepository struct {
	db *pgxpool.Pool
}

func NewURLRepository(db *pgxpool.Pool) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

// create URL
func (r *URLRepository) CreateURL(ctx context.Context, shortURL string, originalURL string) (*model.URL, error) {
	var urlModel model.URL

	err := r.db.QueryRow(ctx, `
		INSERT INTO urls (short_url, original_url)
		VALUES ($1, $2)
		RETURNING id, short_url, original_url, created_at, expires_at, clicks_count
	`, shortURL, originalURL).Scan(&urlModel.ID, &urlModel.ShortURL, &urlModel.OriginalURL, &urlModel.CreatedAt, &urlModel.ExpiresAt, &urlModel.ClickCount)
	if err != nil {
		return nil, err
	}

	return &urlModel, nil
}

// get URL
func (r *URLRepository) GetURL(ctx context.Context, shortURL string) (*model.URL, error) {
	var urlModel model.URL

	err := r.db.QueryRow(ctx, `
		SELECT id, short_url, original_url, created_at, expires_at, clicks_count
		FROM urls
		WHERE short_url = $1
	`, shortURL).Scan(&urlModel.ID, &urlModel.ShortURL, &urlModel.OriginalURL, &urlModel.CreatedAt, &urlModel.ExpiresAt, &urlModel.ClickCount)

	if err != nil {
		return nil, err
	}

	return &urlModel, nil
}
