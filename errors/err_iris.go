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
