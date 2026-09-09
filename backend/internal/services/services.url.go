package services

import (
	"context"
	"errors"
	"goBitly/internal/apperrors"
	"goBitly/internal/model"
	"goBitly/internal/repository"
	"goBitly/internal/utils"
	"log"
)

type URLService struct {
	repo repository.URLRepository
}

func NewURLService(repo repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (url *URLService) CreateShortURLService(ctx context.Context, userId int64, originalURL string) (*model.URL, error) {
	if !utils.IsURLValidUtil(originalURL) {
		return nil, errors.New("invalid URL")
	}

	normalizedURL, err := utils.NormalizeURLUtil(originalURL)
	if err != nil {
		return nil, err
	}

	existingURL, err := url.repo.GetURLByLongURLRepo(ctx, normalizedURL)
	if err != nil && !errors.Is(err, apperrors.ErrURLNotFound) {
		return nil, err
	}

	if existingURL != nil {
		return existingURL, nil
	}
	
	return url.repo.CreateURLRepo(ctx, userId, normalizedURL)
}

func (url *URLService) GetURLByShortURLService(ctx context.Context, shortURL string) (*model.URL, error) {
	existingURL, err := url.repo.GetURLByShortURLRepo(ctx, shortURL)
	if err != nil && !errors.Is(err, apperrors.ErrURLNotFound){
		return nil, err
	}

	return existingURL, nil
}

func (url *URLService) GetURLByOriginalURLService(ctx context.Context, originalURL string) (*model.URL, error) {
	if !utils.IsURLValidUtil(originalURL) {
		return nil, errors.New("invalid URL")
	}

	normalizedURL, err := utils.NormalizeURLUtil(originalURL)
	if err != nil {
		return nil, err
	}

	existingURL, err := url.repo.GetURLByLongURLRepo(ctx, normalizedURL)
	if err != nil && !errors.Is(err, apperrors.ErrURLNotFound) {
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

	return success, nil
}

func (url *URLService) IncrementClickCountService(ctx context.Context, shortURL string) (bool, error) {
	success, err := url.repo.IncrementClickCountRepo(ctx, shortURL)
	if err != nil {
		log.Println("Increment Clicks Failed")
		return false, err
	}

	return success, nil
}

func (url *URLService) GetClickCountsService(ctx context.Context, shortURL string) (int64, error) {
	count, err := url.repo.GetClickCountRepo(ctx, shortURL)
	if err != nil {
		return 0, err
	}

	return count, nil
}
