package xerr

import (
	"fmt"
	"net/http"

	"github.com/file_server/common/constant"
)

/**
常用通用固定错误
*/

type CodeError struct {
	errCode uint32
	errMsg  string
}

// 返回给前端的错误码
func (e *CodeError) GetErrCode() uint32 {
	return e.errCode
}

var requestLanguage = map[string]string{
	"ch": "zh_CN",
	"tw": "zh_TW",
	"en": "en_US",
	"vi": "vi_VN",
	"ru": "ru_RU",
	"ko": "ko_KR",
	"ja": "ja_JP",
	"th": "th_TH",
	"id": "id_ID",
	"tr": "tr_TR",
	"fr": "fr_FR",
	"ar": "ar_AR",
	"es": "es_ES",
	"fa": "fa_FA",
}

// 返回给前端显示端错误信息
func (e *CodeError) GetErrMsg() string {
	return e.errMsg
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrCode:%d，ErrMsg:%s", e.errCode, e.errMsg)
}

func NewErr(code uint32, err error) *CodeError {
	return &CodeError{errCode: code, errMsg: fmt.Sprintf("%v", err)}
}

func NewRateLimitErr(r *http.Request) *CodeError {
	language := r.Header.Get("language")
	msg := "The request frequency is too fast, please try again later"

	var errMsgMap = map[string]string{
		"zh_CN": "请求频率太快，请稍后再试",
		"zh_TW": "請求頻率太快，請稍後再試",
		"fr_FR": "فرکانس درخواست خیلی سریع است ، لطفاً بعداً دوباره امتحان کنید",
		"vi_VN": "Tần số yêu cầu quá nhanh, vui lòng thử lại sau",
		"id_ID": "Minta frekuensi terlalu cepat. Silakan coba lagi nanti",
		"ar_AR": "تردد الطلب سريع جدًا ، يرجى المحاولة مرة أخرى لاحقًا",
		"en_US": "The request frequency is too fast, please try again later",
	}

	v, ok := errMsgMap[language]
	if ok {
		msg = v
	}

	return &CodeError{errCode: constant.REQUEST_RATE_LIMIT_ERROR, errMsg: msg}
}

func NewParamsErr(language string) *CodeError {
	msg := "The verification content is too long and cannot inquire"

	var errMsgMap = map[string]string{
		"zh_CN": "验证内容过长，无法查询",
		"zh_TW": "驗證內容過長，無法查詢",
		"fr_FR": "محتوای تأیید خیلی طولانی است و نمی تواند پرس و جو کند",
		"vi_VN": "Nội dung xác minh quá dài và không thể hỏi",
		"id_ID": "Konten verifikasi terlalu lama dan tidak dapat menanyakan",
		"ar_AR": "محتوى التحقق طويل جدًا ولا يمكن الاستفسار",
		"en_US": "The verification content is too long and cannot inquire",
	}

	v, ok := errMsgMap[language]
	if ok {
		msg = v
	}

	return &CodeError{errCode: constant.REQUEST_PARAM_ERROR, errMsg: msg}
}
