package domain

import "context"

// Usecase represent the article's usecases
type Usecase interface {
	GetArticle(ctx context.Context, id *string) (*Article, error)
	AddArticle(ctx context.Context, id *string) (*Article, error)
	EditArticle(ctx context.Context, id *string) (*Article, error)
	DeleteArticle(ctx context.Context, id *string) (*Article, error)
}

// ITonTon represent the repository contract
type ITonTon interface {
	GetArticle(ctx context.Context, id *string) (*Article, error)
	AddArticle(ctx context.Context, id *string) (*Article, error)
	EditArticle(ctx context.Context, id *string) (*Article, error)
	DeleteArticle(ctx context.Context, id *string) (*Article, error)
}
