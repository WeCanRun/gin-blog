package dto

type GetListRequest struct {
	PageNum  uint `json:"page_num" form:"page_num" binding:"gte=1,required"`
	PageSize uint `json:"page_size" form:"page_size" binding:"gte=1,required"`
}

type Pager struct {
	PageNum   uint `json:"page_num"`
	PageSize  uint `json:"page_size"`
	TotalRows uint `json:"total_rows"`
}
