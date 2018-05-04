package errors

type IrisError struct {
	ErrCode string
	ErrMsg  string
}

func (e *IrisError) Error() string {
	return e.ErrMsg
}

func (e *IrisError) IsNotNull() bool {
	if e.ErrCode != "" {
		return true
	}
	return false
}

func (e *IrisError) New(errCode string, errMsg string) IrisError {
	return IrisError{
		ErrCode: errCode,
		ErrMsg:  errMsg,
	}
}
