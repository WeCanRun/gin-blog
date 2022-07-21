package errcode

const (
	SUCCESS        = 200
	INVALID_PARAMS = 400
	SERVER_ERROR   = 500

	ERROR                   = 502
	ERROR_EXIST_TAG         = 10001
	ERROR_NOT_EXIST_TAG     = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUTH_CHECK_TOKEN_FAIL    = 20001
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT = 20002
	ERROR_AUTH_TOKEN               = 20003
	ERROR_AUTH                     = 20004
)

type Response struct {
	Status int         `json:"status" example:"200" example:"400" example:"500" example:"502"`
	Msg    string      `json:"msg" example:"ok"`
	Data   interface{} `json:"data,omitempty"`
}
