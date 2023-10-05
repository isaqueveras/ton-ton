package domain

import "time"

const (
	// ArticlePublishStatus indicates the article is as published
	ArticlePublishStatus = "Publish"
)

// Article is representing the Article data struct
type Article struct {
	ID        *string    `json:"id,omitempty"`
	Title     *string    `json:"title,omitempty"`
	Content   *string    `json:"content,omitempty"`
	Status    *string    `json:"status,omitempty"`
	Slug      *string    `json:"slug,omitempty"`
	CreatorID *string    `json:"creator_id,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
