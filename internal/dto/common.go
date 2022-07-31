package dto

type GetListRequest struct {
	PageNum  uint `json:"page_num"`
	PageSize uint `json:"page_size"`
}

type Pager struct {
	PageNum   uint `json:"page_num"`
	PageSize  uint `json:"page_size"`
	TotalRows uint `json:"total_rows"`
}
