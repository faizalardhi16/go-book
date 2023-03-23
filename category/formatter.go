package category

type CategoryResponse struct {
	ID           string `json:"id"`
	CategoryName string `json:"categoryName"`
}

func FormatCategory(category Category) CategoryResponse {
	categoryResponse := CategoryResponse{}

	categoryResponse.ID = category.ID
	categoryResponse.CategoryName = category.CategoryName

	return categoryResponse
}
