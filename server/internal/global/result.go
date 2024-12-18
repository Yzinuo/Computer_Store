package global

var(
	SUCCESS = 0
	FAIL = 500
	_code  = map[int]struct{}{}
	_msg   = make(map[int]string)
)

type Result struct{
	Code int 
	Msg  string
}

func RegisterErrorCode(code int, msg string) (Result)  {
	if _,ok := _code[code]; ok {
		panic("code has been registered")
	}
	if msg == ""{
		panic("msg cannot be empty")
	}

	_code[code] = struct{}{}

	_msg[code] = msg

	return Result{
		Code: code,
		Msg: msg,
	}
}

func GetMsg(code int) string{
	return _msg[code]
}


var(
	OkReresult = RegisterErrorCode(SUCCESS, "ok")
	FailResult = RegisterErrorCode(FAIL, "fail")

	ErrRequest  = RegisterErrorCode(9001, "请求参数格式错误")
	ErrDbOp     = RegisterErrorCode(9004, "数据库操作异常")
	ErrRedisOp  = RegisterErrorCode(9005, "Redis 操作异常")

	ErrPassword     = RegisterErrorCode(1002, "密码错误")
	ErrUserNotExist = RegisterErrorCode(1003, "该用户不存在")

	ErrTokenNotExist    = RegisterErrorCode(1201, "TOKEN 不存在，请重新登陆")
	ErrTokenRuntime     = RegisterErrorCode(1202, "TOKEN 已过期，请重新登陆")
	ErrTokenWrong       = RegisterErrorCode(1203, "TOKEN 不正确，请重新登陆")
	ErrTokenType        = RegisterErrorCode(1204, "TOKEN 格式错误，请重新登陆")
	ErrTokenCreate      = RegisterErrorCode(1205, "TOKEN 生成失败")

	ErrSendEmail = RegisterErrorCode(6101, "发送邮件失败")
	ErrCodeNoexit = RegisterErrorCode(6102, "Code不存在 请重新注册")
	ErrParseEmailCode = RegisterErrorCode(6103, "解析邮件Code失败 请重试")
	ErrUserExist = RegisterErrorCode(6104, "该邮箱已经注册 请重新注册")

		// 上传或获取文件
	ErrFileReceive = RegisterErrorCode(9101, "文件接收失败")
	ErrFileUpload  = RegisterErrorCode(9100, "文件上传失败")
)