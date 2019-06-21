package util

var errorCode = map[int]string{
	0:    "ok",
	-1:   "系统异常，请稍后重试",
	-2:   "输入数据有误",
	-3:   "请检查账号和密码",
	-4:   "请输入正确的手机号码",
	1201: "不允许删除",
}

//ErrorCode 错误代码描述
func ErrorCode(code int) string {
	var codeString = ""
	if _, ok := errorCode[code]; ok {
		codeString = errorCode[code]
	}
	return codeString
}
