CREATE TYPE article_status AS ENUM ('Draft', 'Pending', 'Private', 'Publish', 'Trash');

CREATE TABLE t_article (
	id VARCHAR(8) PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
  content TEXT,
	creator_id VARCHAR(8),
	"status" article_status DEFAULT 'Draft',
	slug VARCHAR(150) NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMPTZ
);
