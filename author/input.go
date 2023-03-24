package author

type AuthorInput struct {
	AuthorName string `json:"authorName" binding:"required"`
}

type DeleteAuthorInput struct {
	ID string `json:"id" binding:"required"`
}
