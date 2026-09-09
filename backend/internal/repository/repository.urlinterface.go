package repository

import (
	"context"
	"goBitly/internal/model"
)

type URLRepository interface {
	CreateURLRepo(ctx context.Context, userId int64, originalURL string) (*model.URL, error)
	GetURLByShortURLRepo(ctx context.Context, shortURL string) (*model.URL, error)
	GetURLByLongURLRepo(ctx context.Context, originalURL string) (*model.URL, error)
	DeleteURLRepo(ctx context.Context, shortURL string) (bool, error)
	IncrementClickCountRepo(ctx context.Context, shortURL string) (bool, error)
	GetClickCountRepo(ctx context.Context, shortURL string) (int64, error)
}
