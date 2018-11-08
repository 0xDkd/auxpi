package e

var MsgFlags = map[int]string {
	SUCCESS : "ok",
	ERROR : "fail",
	INVALID_PARAMS : "请求参数错误",
	ERROR_FILE_IS_EMPTY : "上传文件为空",
	ERROR_FILE_IS_TOO_LARGE : "上传文件太大",
	ERROR_FILE_TYPE: "文件类型错误",
	ERROR_CAN_NOT_GET_IMG_URL : "无法获取第三方图床 URL",
	ERROR_TOO_MANY_IMAGES : "上传图片太多",
	ERROR_AUTH_CHECK_TOKEN_FAIL : "Token鉴权失败",
	ERROR_ACCESS_DENIED : "禁止访问",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
