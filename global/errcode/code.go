package errcode

var (
	Success     = New(0, "ok")
	BadRequest  = New(400, "请求参数错误")
	ServerError = New(500, "系统内部错误")
	BadGateway  = New(502, "请求异常")

	ExistTag             = New(10001, "已存在该标签名称")
	ErrorNotExistTag     = New(10002, "该标签不存在")
	ErrorNotExistArticle = New(10003, "该文章不存在")

	ErrorAuthCheckTokenFail    = New(20001, "Token鉴权失败")
	ErrorAuthCheckTokenTimeout = New(20002, "Token已超时")
	ErrorAuthToken             = New(20003, "Token生成失败")
	ErrorAuth                  = New(20004, "Token错误")
)
