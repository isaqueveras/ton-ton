package repository

import (
	"context"
	"database/sql"
	"ton-ton/domain"
)

type repository struct {
	pg *sql.DB
}

func New(conn *sql.DB) domain.ITonTon {
	return &repository{pg: conn}
}

func (r *repository) GetArticle(ctx context.Context, id *string) (*domain.Article, error) {
	res := new(domain.Article)

	query := `SELECT title, content, status, creator_id, slug, created_at, updated_at FROM t_article WHERE id = $1`
	row := r.pg.QueryRowContext(ctx, query, id)

	if err := row.Scan(&res.Title, &res.Content, &res.Status, &res.CreatorID, &res.Slug, &res.CreatedAt, &res.UpdatedAt); err != nil {
		return nil, err
	}

	return res, nil
}

func (r *repository) AddArticle(ctx context.Context, article *domain.Article) error {
	_, err := r.pg.ExecContext(ctx, "INSERT INTO t_article (id, title, content, status, creator_id, slug) VALUES ($1, $2, $3, $4, $5, $6)",
		article.ID, article.Title, article.Content, article.Status, article.CreatorID, article.Slug)
	return err
}

func (r *repository) EditArticle(ctx context.Context, id *string) (*domain.Article, error) {
	return nil, nil
}

func (r *repository) DeleteArticle(ctx context.Context, id *string) (*domain.Article, error) {
	return nil, nil
}
