package errors

type IrisError struct {
	ErrCode uint32
	ErrMsg  string
}

func (e IrisError) Error() string {
	return e.ErrMsg
}

func (e IrisError) IsNotNull() bool {
	if e.ErrCode != 0 {
		return true
	}
	return false
}

func (e IrisError) New(errCode uint32, errMsg string) IrisError {
	return IrisError{
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}

func SdkCodeToIrisErr(code uint32, msg string) IrisError {
	errFun, ok := sdkCodeToErrFunc[CodeType(code)]
	if ok {
		return errFun(msg)
	}
	return ExtSysUnKnownErr("not existed code define,code:%d", code)
}
