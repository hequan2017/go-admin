package e

var MsgFlags = map[int]string{
	SUCCESS:        "ok",
	ERROR:          "fail",
	INVALID_PARAMS: "请求参数错误",

	ERROR_EXIST:       "已存在该对象名称",
	ERROR_EXIST_FAIL:  "获取已存在对象失败",
	ERROR_NOT_EXIST:   "该对象不存在",
	ERROR_GET_S_FAIL:  "获取所有对象失败",
	ERROR_COUNT_FAIL:  "统计对象失败",
	ERROR_ADD_FAIL:    "新增对象失败",
	ERROR_EDIT_FAIL:   "修改对象失败",
	ERROR_DELETE_FAIL: "删除对象失败",
	ERROR_EXPORT_FAIL: "导出对象失败",
	ERROR_IMPORT_FAIL: "导入对象失败",

	ERROR_AUTH_CHECK_TOKEN_FAIL:    "Token鉴权失败",
	ERROR_AUTH_CHECK_TOKEN_TIMEOUT: "Token已超时",
	ERROR_AUTH_TOKEN:               "Token生成失败",
	ERROR_AUTH:                     "Token错误",
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	}

	return MsgFlags[ERROR]
}
