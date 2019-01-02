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

func IrisErr(errCode uint32, errMsg string) func(errMsg error) IrisError {
	return func(err error) IrisError {
		msg := errMsg
		if err != nil {
			msg = errMsg + err.Error()
		}
		return IrisError{
			ErrCode: errCode,
			ErrMsg:  msg,
		}
	}
}

func SdkCodeToIrisErr(code uint16) IrisError {
	return sdkCodeToIrisCodeMap[code]
}
