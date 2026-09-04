package services

import "goBitly/internal/repository"

type URLService struct {
	repo *repository.URLRepository
}

func NewURLService(repo *repository.URLRepository) *URLService {
	return &URLService{
		repo: repo,
	}
}

func (url *URLService) CreateShortURLService(originalURL string) (string, error) {

}

func (url *URLService) GetOriginalURLService() {

}

func (url *URLService) DeleteURLService() {

}

func (url *URLService) IncrementClickCountService() {

}
