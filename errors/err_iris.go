package errors

import "fmt"

type CodeType uint32

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

func SdkCodeToIrisErr(space string, code uint16, msg string) IrisError {
	c := sdkCode(space, CodeType(code))
	errFun, ok := sdkCodeToErrFunc[c]
	if ok {
		return errFun(msg)
	}
	return UnKnownErr("not existed code difine,space:%s,code:%d", space, code)
}

func sdkCode(space string, code CodeType) string {
	return fmt.Sprintf("%s-%d", space, code)
}
