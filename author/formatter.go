package author

type FormatAuthorResponse struct {
	AuthorName string `json:"authorName"`
}

func FormatterAuthor(author Author) FormatAuthorResponse {
	res := FormatAuthorResponse{}

	res.AuthorName = author.AuthorName

	return res
}
