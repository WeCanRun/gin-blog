package dto

type GetTagsRequest struct {
	GetListRequest
}

type GetTagResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type GetTagsResponse struct {
	IDs   []uint   `json:"ids"`
	Names []string `json:"names"`
}

type AddTagRequest struct {
	Name      string `json:"name" binding:"required"`
	CreatedBy string `json:"created_by"`
	State     int    `json:"state"`
}

type EditRequest struct {
	ID        uint   `json:"id" binding:"required"`
	Name      string `json:"name"`
	UpdatedBy string `json:"update_by"`
	State     int    `json:"state"`
}

type ExportTagsRequest struct {
	Name  string `json:"name"`
	State int    `json:"state"`
}

type ExportTagsResponse struct {
	ExportFullUrl string `json:"export_full_url"`
	ExportSaveUrl string `json:"export_save_url"`
}
