package services

import (
	"context"
	"errors"
	"goBitly/internal/model"
	"goBitly/internal/repository"
	"goBitly/internal/utils"

	"github.com/jackc/pgx/v5"
)

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (url *URLService) CreateShortURLService(ctx context.Context, originalURL string) (*model.URL, error) {
	if originalURL == ""{
		return nil, errors.New("url can't be empty")
	}
	
	if !utils.IsURLValidUtil(originalURL) {
		return nil, errors.New("invalid URL")
	}

	normalizedURL, err := utils.NormalizeURLUtil(originalURL)
	if err != nil {
		return nil, err
	}

	existingURL, err := url.repo.GetURLByLongURLRepo(ctx, normalizedURL)
	if err != nil && !errors.Is(err, pgx.ErrNoRows){
		return nil, err
	}
	if existingURL != nil {
		return existingURL, nil
	}

	shortURL := utils.ShortURLGenerator(5)
	urlModel, err := url.repo.CreateURLRepo(ctx, shortURL, normalizedURL)
	if err != nil {
		return nil, err
	}

	return urlModel, nil
}

func (url *URLService) GetURLByShortURLService(ctx context.Context, shortURL string) (*model.URL, error){
	if shortURL == "" {
		return nil, errors.New("short URL can't be empty")
	}

	existingURL, err := url.repo.GetURLByShortURLRepo(ctx, shortURL)
	if err != nil {
		return nil, err
	}
	
	return existingURL, nil
}

func (url *URLService) GetURLByOriginalURLService(ctx context.Context, originalURL string) (*model.URL, error){
	if !utils.IsURLValidUtil(originalURL) {
		return nil, errors.New("invalid URL")
	}

	normalizedURL, err := utils.NormalizeURLUtil(originalURL)
	if err != nil {
		return nil, err
	}
	
	existingURL, err := url.repo.GetURLByLongURLRepo(ctx, normalizedURL)
	if err != nil && !errors.Is(err, pgx.ErrNoRows){
		return nil, err
	}
	if existingURL != nil {
		return existingURL, nil
	}
	
	return nil, err
}

func (url *URLService) DeleteURLService(ctx context.Context, shortURL string) (bool, error) {
	success, err := url.repo.DeleteURLRepo(ctx, shortURL)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}

	return true, nil
}

func (url *URLService) IncrementClickCountService(ctx context.Context, shortURL string) (bool, error) {
	success, err := url.repo.IncrementClickCountsRespo(ctx, shortURL)
	if err != nil {
		return false, err
	}
	if !success {
		return false, nil
	}

	return true, nil
}

func (url* URLService) GetClickCountsService(ctx context.Context, shortURL string) (int64, error) {
	count, err := url.repo.GetClickCountsRepo(ctx, shortURL)
	if err != nil {
		return -1, err
	}
	if count <=1 {
		return -1, nil
	}

	return count, nil
}

