package resx

import "constantx"

type Res struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	ErrData interface{} `json:"err_data"`
}

type PageList struct {
	Total int64       `json:"total"`
	List  interface{} `json:"list"`
}

/**
正确返回
*/
func ResSuccess(errData interface{}) *Res {
	return &Res{ErrCode: constantx.ResSuccessCode, ErrMsg: constantx.ResSuccess, ErrData: errData}
}

/**
错误返回
*/
func ResFail(errMsg string) *Res {
	return &Res{ErrCode: constantx.ResFailCode, ErrMsg: errMsg, ErrData: nil}
}

/**
返回列表
*/
func ResPageList(total int64, list interface{}) *Res {
	pageList := &PageList{total, list}
	return ResSuccess(pageList)
}
