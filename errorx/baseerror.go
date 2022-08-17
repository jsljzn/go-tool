package errorx

const defaultCode = 4005

type CodeError struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	ErrData interface{} `json:"err_data"`
}

type CodeErrorResponse struct {
	ErrCode int         `json:"err_code"`
	ErrMsg  string      `json:"err_msg"`
	ErrData interface{} `json:"err_data"`
}

func NewCodeError(errCode int, errMsg string) error {
	return &CodeError{ErrCode: errCode, ErrMsg: errMsg}
}

func NewDefaultError(msg string) error {
	return NewCodeError(defaultCode, msg)
}

func (e *CodeError) Error() string {
	return e.ErrMsg
}

func (e *CodeError) Data() *CodeErrorResponse {
	return &CodeErrorResponse{
		ErrCode: e.ErrCode,
		ErrMsg:  e.ErrMsg,
	}
}
