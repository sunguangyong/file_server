package constant

const (
	OK                       uint32 = 200    //成功返回
	REPEATVERIFY             uint32 = 201    //认证信息数据重复
	SERVER_COMMON_ERROR      uint32 = 100001 //全局通用错误
	REQUEST_PARAM_ERROR      uint32 = 100002 //请求参数错误
	REQUEST_RATE_LIMIT_ERROR uint32 = 100003 //触发限流
	REQUEST_SIGN_ERROR       uint32 = 100004 //签名错误
)
