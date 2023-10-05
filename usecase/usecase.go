package usecase

import (
	"context"
	"time"
	"ton-ton/domain"
)

type usecase struct {
	repo    domain.ITonTon
	timeout time.Duration
}

func NewUsecase(a domain.ITonTon, timeout time.Duration) domain.Usecase {
	return &usecase{repo: a, timeout: timeout}
}

func (a *usecase) GetArticle(ctx context.Context, id *string) (article *domain.Article, err error) {
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()
	return a.repo.GetArticle(ctx, id)
}

func (a *usecase) AddArticle(ctx context.Context, id *string) (*domain.Article, error) {
	return nil, nil
}

func (a *usecase) EditArticle(ctx context.Context, id *string) (*domain.Article, error) {
	return nil, nil
}

func (a *usecase) DeleteArticle(ctx context.Context, id *string) (*domain.Article, error) {
	return nil, nil
}
