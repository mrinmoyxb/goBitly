package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"goBitly/internal/model"
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
func (r *URLRepository) CreateURLRepo(ctx context.Context, shortURL string, originalURL string) (*model.URL, error) {
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
func (r *URLRepository) GetURLByShortURLRepo(ctx context.Context, shortURL string) (*model.URL, error) {
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

// get URL by original URL
func (r *URLRepository) GetURLByLongURLRepo(ctx context.Context, longURL string) (*model.URL, error) {
	var urlModel model.URL

	err := r.db.QueryRow(ctx, `
		SELECT id, short_url, original_url, created_at, expires_at, clicks_count
		FROM urls
		WHERE original_url = $1
	`, longURL).Scan(&urlModel.ID, &urlModel.ShortURL, &urlModel.OriginalURL, &urlModel.CreatedAt, &urlModel.ExpiresAt, &urlModel.ClickCount)
	if err != nil {
		return nil, err
	}

	return &urlModel, nil
}

// get URL by short url
func (r *URLRepository) ExistsShortURLRepo(ctx context.Context, shortURL string) (bool, error) {
	var exists bool

	err := r.db.QueryRow(ctx, `
		SELECT EXISTS(
			SELECT 1 FROM urls where short_url = $1
		)
	`, shortURL).Scan(&exists)
	if err != nil {
		return false, err
	}

	return exists, nil
}

// delete URL
func (r *URLRepository) DeleteURLRepo(ctx context.Context, shortURL string) (bool, error) {
	cmdTag, err := r.db.Exec(ctx, `
	DELETE FROM urls WHERE short_url = $1
	`, shortURL)

	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}

// increment count
func (r *URLRepository) IncrementClickCountsRespo(ctx context.Context, shortURL string) (bool, error) {
	cmdTag, err := r.db.Exec(ctx, `
		UPDATE urls SET clicks_count = clicks_count + 1 
		WHERE short_url = $1
	`, shortURL)

	if err != nil {
		return false, err
	}

	if cmdTag.RowsAffected() == 0 {
		return false, nil
	}

	return true, nil
}

// clicks count
func (r *URLRepository) GetClickCountsRepo(ctx context.Context, shortURL string) (int64, error) {
	var ClickCountModel model.ClickCount

	err := r.db.QueryRow(ctx, `
		SELECT clicks_count FROM urls WHERE short_url = $1
	`, shortURL).Scan(&ClickCountModel.ClickCount)
	if err != nil {
		return -1, err
	}

	return ClickCountModel.ClickCount, nil
}
