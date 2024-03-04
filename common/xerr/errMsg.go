package xerr

import "github.com/file_server/common/constant"

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[constant.OK] = "success"
	message[constant.SERVER_COMMON_ERROR] = "server error"
	message[constant.REQUEST_SIGN_ERROR] = "signature error"
}

func MapErrMsg(errcode uint32) string {
	if msg, ok := message[errcode]; ok {
		return msg
	} else {
		return "server error"
	}
}

func IsCodeErr(errcode uint32) bool {
	if _, ok := message[errcode]; ok {
		return true
	} else {
		return false
	}
}
