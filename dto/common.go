package dto

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

type GetListRequest struct {
	PageNum  uint `json:"page_num"`
	PageSize uint `json:"page_size"`
}

