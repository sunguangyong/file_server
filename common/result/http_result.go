package result

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/file_server/common/constant"

	"encoding/json"

	"github.com/file_server/common/xerr"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"google.golang.org/grpc/status"
)

func HttpResult(r *http.Request, w http.ResponseWriter, req, resp interface{}, err error) {
	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := constant.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
			httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
			return
		}
		if e, ok := causeErr.(*xerr.CtxErr); ok {
			reqBytes, _ := json.Marshal(req)
			responseBytes, _ := json.Marshal(resp)
			logx.WithContext(r.Context()).Errorf("【API-CTX-ERR】Request:%v,Response:%v,ctx:%v,err:%v,stack:%v", string(reqBytes), string(responseBytes), e.Content, err, string(debug.Stack()))
			httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
			return
		}

		if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
			grpcCode := uint32(gstatus.Code())
			if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
				errcode = grpcCode
				errmsg = gstatus.Message()
			}
			httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
			return
		}
		reqBytes, _ := json.Marshal(req)
		responseBytes, _ := json.Marshal(resp)
		logx.WithContext(r.Context()).Errorf("【API-UNKNOW-ERR】Request:%v,Response:%v,err:%v,stack:%v", string(reqBytes), string(responseBytes), err, string(debug.Stack()))
		httpx.WriteJson(w, http.StatusBadRequest, Error(errcode, errmsg))
		return
	}
}

// 授权的http方法
func AuthHttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {

	if err == nil {
		//成功返回
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		//错误返回
		errcode := constant.SERVER_COMMON_ERROR
		errmsg := "服务器开小差啦，稍后再来试一试"

		causeErr := errors.Cause(err)                // err类型
		if e, ok := causeErr.(*xerr.CodeError); ok { //自定义错误类型
			//自定义CodeError
			errcode = e.GetErrCode()
			errmsg = e.GetErrMsg()
		} else {
			if gstatus, ok := status.FromError(causeErr); ok { // grpc err错误
				grpcCode := uint32(gstatus.Code())
				if xerr.IsCodeErr(grpcCode) { //区分自定义错误跟系统底层、db等错误，底层、db错误不能返回给前端
					errcode = grpcCode
					errmsg = gstatus.Message()
				}
			}
		}

		logx.WithContext(r.Context()).Errorf("【GATEWAY-ERR】 : %+v ", err)
		httpx.WriteJson(w, http.StatusUnauthorized, Error(errcode, errmsg))
	}
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	errMsg := fmt.Sprintf("%s ,%s", xerr.MapErrMsg(constant.REQUEST_PARAM_ERROR), err.Error())
	httpx.WriteJson(w, http.StatusBadRequest, Error(constant.REQUEST_PARAM_ERROR, errMsg))
}
