package dto

type GetArticlesRequest struct {
	GetListRequest
}

type GetArticlesResponse struct {
	IDs      []uint   `json:"ids"`
	Titles   []string `json:"titles"`
	TagIds   []uint   `json:"tag_ids"`
	TagNames []string `json:"tag_names"`
}

type commonRequest struct {
	TagId   uint   `json:"tag_id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	State   int    `json:"state"`
}

type AddArticleRequest struct {
	commonRequest
	CreatedBy     string `json:"created_by"`
	CoverImageUrl string `json:"cover_image_url"`
}

type EditArticleRequest struct {
	ID uint `json:"id" binding:"required"`
	commonRequest
	CoverImageUrl string `json:"cover_image_url"`
	UpdatedBy     string `json:"updated_by"`
}

type GenArticlePosterReq struct {
	Width  int `json:"width" binding:"required"`
	Height int `json:"height" binding:"required"`
}

type GenArticlePosterResp struct {
	PosterSaveUrl string `json:"poster_save_url"`
	PosterUrl     string `json:"poster_url"`
}
