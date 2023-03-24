package author

import "time"

type Author struct {
	ID         string
	AuthorName string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
