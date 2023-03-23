package category

type CategoryInput struct {
	CategoryName string `json:"categoryName" binding:"required"`
}
