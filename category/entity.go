package category

import "time"

type Category struct {
	ID           string `json:"categoryId"`
	CategoryName string `json:"categoryName"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
