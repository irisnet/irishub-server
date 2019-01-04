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

func SdkCodeToIrisErr(space string, code uint16) IrisError {
	c := sdkCode(space, CodeType(code))
	err, ok := sdkCodeToIrisCodeMap[c]
	if ok {
		return err
	}
	return UnKnownErr(fmt.Errorf("not existed code difine,space:%s,code:%d", space, code))
}

func sdkCode(space string, code CodeType) string {
	return fmt.Sprintf("%s-%d", space, code)
}
