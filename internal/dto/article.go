package dto

type GetArticlesRequest struct {
	GetListRequest
}

type GetArticlesResponse struct {
	Pager    Pager    `json:"pager"`
	IDs      []uint   `json:"ids"`
	Titles   []string `json:"titles"`
	TagIds   []uint   `json:"tag_ids"`
	TagNames []string `json:"tag_names"`
}

type commonRequest struct {
	TagId   uint   `json:"tag_id" binding:"required"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
	State   int    `json:"state" binding:"required,oneof=0 1"`
}

type AddArticleRequest struct {
	commonRequest
	CreatedBy     string `json:"created_by",binding:"required, min=3, max=100"`
	CoverImageUrl string `json:"cover_image_url"`
}

type EditArticleRequest struct {
	ID uint `json:"id" binding:"required,gte=1"`
	commonRequest
	CoverImageUrl string `json:"cover_image_url"`
	UpdatedBy     string `json:"updated_by",binding:"required,min=3,max=100"`
}

type GenArticlePosterReq struct {
	Width  int `json:"width" binding:"required,gte=33"`
	Height int `json:"height" binding:"required,gte=33"`
}

type GenArticlePosterResp struct {
	PosterSaveUrl string `json:"poster_save_url"`
	PosterUrl     string `json:"poster_url"`
}
