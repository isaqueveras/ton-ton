package domain

import "context"

// Usecase represent the article's usecases
type Usecase interface {
	GetArticle(context.Context, *string) (*Article, error)
	AddArticle(context.Context, *Article) error
	EditArticle(context.Context, *Article) error
	DeleteArticle(context.Context, *string) error
}

// ITonTon represent the repository contract
type ITonTon interface {
	GetArticle(context.Context, *string) (*Article, error)
	AddArticle(context.Context, *Article) error
	EditArticle(context.Context, *Article) error
	DeleteArticle(context.Context, *string) error
}
