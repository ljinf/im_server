package v1

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")

	// more biz errors
	ErrCreateIdFailed       = newError(10001, "申请用户ID失败")
	ErrCreateUserFailed     = newError(10002, "创建用户失败")
	ErrCreateUserInfoFailed = newError(10003, "初始化用户消息失败")
	ErrEmailAlreadyUse      = newError(10004, "邮件已被注册")
	ErrPhoneAlreadyUse      = newError(10005, "手机号码已被注册")
	ErrUpdateUserInfoFailed = newError(10006, "更新信息失败")
)
