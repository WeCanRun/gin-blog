package dto

type GetTagsRequest struct {
	GetListRequest
}

type GetTagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetTagsResponse struct {
	Pager Pager    `json:"pager"`
	IDs   []uint   `json:"ids"`
	Names []string `json:"names"`
}

type AddTagRequest struct {
	Name      string `json:"name" binding:"required, min=4, max=16" minLength:"4" maxLength:"16" example:"random string"`
	CreatedBy string `json:"created_by" binding:"required, min=3, max=100"`
	State     int    `json:"state" binding:"oneof=0 1"`
}

type EditRequest struct {
	ID        uint   `json:"id" binding:"required,gte=1"`
	Name      string `json:"name" binding:"required, min=4, max=16"`
	UpdatedBy string `json:"update_by" binding:"required, min=3, max=100"`
	State     int    `json:"state" binding:"oneof=0 1"`
}

type ExportTagsRequest struct {
	Name  string `json:"name" binding:"required,min=3,max=100"`
	State int    `json:"state" binding:"oneof=0 1"`
}

type ExportTagsResponse struct {
	ExportFullUrl string `json:"export_full_url"`
	ExportSaveUrl string `json:"export_save_url"`
}
