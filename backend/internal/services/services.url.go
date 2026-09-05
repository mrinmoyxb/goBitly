package services

import (
	"context"
	"errors"
	"goBitly/internal/model"
	"goBitly/internal/repository"
	"goBitly/internal/utils"
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
	if utils.IsURLValidUtil(originalURL) {
		return nil, errors.New("invalid URL")
	}

	normalizedURL, err := utils.NormalizeURLUtil(originalURL)
	if err != nil {
		return nil, err
	}

	existingURL, err := url.repo.GetURLByLongURLRepo(ctx, normalizedURL)
	if err != nil {
		return nil, err
	}

	if existingURL != nil {
		return existingURL, nil
	}

	return nil, nil
}

func (url *URLService) GetOriginalURLService() {

}

func (url *URLService) DeleteURLService() {

}

func (url *URLService) IncrementClickCountService() {

}
