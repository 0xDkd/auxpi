package e

var MsgFlags = map[int]string{
	SUCCESS: "ok",
	ERROR:   "fail",

	INVALID_PARAMS: "请求参数错误",

	ERROR_FILE_IS_EMPTY:            "上传文件为空",
	ERROR_FILE_IS_TOO_LARGE:        "上传文件太大",
	ERROR_FILE_TYPE:                "文件类型错误",
	ERROR_CAN_NOT_GET_IMG_URL:      "无法获取第三方图床 URL",
	ERROR_TOO_MANY_IMAGES:          "上传图片太多",
	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token 过期",
	ERROR_AUTH_TOKEN:               "Token 不正确",
	ERROR_AUTH:                     "认证失败",
	ERROR_ACCESS_DENIED:            "禁止访问",

	ERROR_TASK_REPEAT: "任务重复提交,请等待当前任务完成",

	ERROR_USER_LOGIN: "用户不存在或用户名密码错误",

	ERROR_USER_COOKIE: "用户 COOKIE 错误",

	ERROR_USER_ALREADY_EXIST: "用户已存在",
	ERROR_USER_NOT_EXIST:     "用户不存在",
	ERROR_USER_RESET_TOKEN:   "重置 Token 错误",
	ERROR_USER_UN_LOGIN:      "用户未登录",

	ERROR_UPLOAD_PARAM:       "上传参数错误",
	ERROR_CAN_NOT_UPLOAD:     "无法上传图片到第三方图床",
	ERROR_UPLOAD_TOKEN_ERROR: "上传 Token 错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}
	return MsgFlags[ERROR]
}
