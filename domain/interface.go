package domain

import "context"

// Usecase represent the article's usecases
type Usecase interface {
	GetArticle(context.Context, *string) (*Article, error)
	AddArticle(context.Context, *Article) error
	EditArticle(context.Context, *string) (*Article, error)
	DeleteArticle(context.Context, *string) (*Article, error)
}

// ITonTon represent the repository contract
type ITonTon interface {
	GetArticle(ctx context.Context, id *string) (*Article, error)
	AddArticle(context.Context, *Article) error
	EditArticle(ctx context.Context, id *string) (*Article, error)
	DeleteArticle(ctx context.Context, id *string) (*Article, error)
}
