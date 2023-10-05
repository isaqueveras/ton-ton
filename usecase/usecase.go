package usecase

import (
	"context"
	"strings"
	"time"
	"ton-ton/domain"
	"ton-ton/utils"

	"github.com/google/uuid"
	"github.com/gosimple/slug"
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

func (a *usecase) AddArticle(ctx context.Context, article *domain.Article) error {
	ctx, cancel := context.WithTimeout(ctx, a.timeout)
	defer cancel()

	id, err := uuid.NewUUID()
	if err != nil {
		return err
	}

	article.ID = utils.Pointer(strings.Split(id.String(), "-")[0])
	article.Slug = utils.Pointer(slug.Make(*article.Title))

	return a.repo.AddArticle(ctx, article)
}

func (a *usecase) EditArticle(ctx context.Context, id *string) (*domain.Article, error) {
	return nil, nil
}

func (a *usecase) DeleteArticle(ctx context.Context, id *string) (*domain.Article, error) {
	return nil, nil
}
