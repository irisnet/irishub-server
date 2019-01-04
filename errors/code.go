package errors

import (
	e "errors"
	"fmt"
)

const (
	IrishubErrMsg = "irishub error:%s"

	EC40001 = 40001
	EC40002 = 40002

	EC50001 = 50001
	EC50002 = 50002

	EC60001 = 60001
	EC60002 = 60002
	EC60003 = 60003
	EC60004 = 60004
	EC60005 = 60005
	EC60006 = 60006
	EC60007 = 60007
)

var (
	InvalidParamsErr = errFun(EC40001, "invalid param error: %s")
	ParamConvertErr  = errFun(EC40002, "param convert error: %s")
	SysErr           = errFun(EC50001, "system error: %s")
	UnKnownErr       = errFun(EC50002, "unKnown error: %s")
	TxExistedErr     = errFun(EC60001, "tx alreay existed error: %s ")
	TxTimeoutErr     = errFun(EC60002, "tx timeout error: %s")
)

//sdkCode
const (
	//commom
	RootCodeSpace                = "sdk"
	CodeInternal        CodeType = 1
	CodeInvalidSequence CodeType = 3
	CodeUnauthorized    CodeType = 4
	CodeUnknownRequest  CodeType = 6
	CodeInvalidAddress  CodeType = 7
	CodeOutOfGas        CodeType = 12

	//stake
	StakeCodeSpace                 = "stake"
	CodeInvalidValidator  CodeType = 101
	CodeInvalidDelegation CodeType = 102
	CodeInvalidInput      CodeType = 103
	CodeValidatorJailed   CodeType = 104

	//distribution
	DistrCodeSpace                  = "distr"
	CodeNoDistributionInfo CodeType = 104
)

var sdkCodeToIrisCodeMap = map[string]IrisError{
	sdkCode(RootCodeSpace, CodeInternal):        errFun(EC50002, IrishubErrMsg)(e.New("internal error")),
	sdkCode(RootCodeSpace, CodeUnauthorized):    errFun(EC50002, IrishubErrMsg)(e.New("unauthorized")),
	sdkCode(RootCodeSpace, CodeUnknownRequest):  errFun(EC50002, IrishubErrMsg)(e.New("unknown request")),
	sdkCode(RootCodeSpace, CodeInvalidAddress):  errFun(EC50002, IrishubErrMsg)(e.New("invalid address")),
	sdkCode(RootCodeSpace, CodeInvalidSequence): errFun(EC50002, IrishubErrMsg)(e.New("Invalid sequence")),
	//distinct
	sdkCode(RootCodeSpace, CodeOutOfGas): errFun(EC60003, IrishubErrMsg)(e.New("out of gas")),

	sdkCode(StakeCodeSpace, CodeInvalidValidator):  errFun(EC50002, IrishubErrMsg)(e.New("validator does not exist for that address")),
	sdkCode(StakeCodeSpace, CodeInvalidDelegation): errFun(EC50002, IrishubErrMsg)(e.New("no delegation for this validator")),
	sdkCode(StakeCodeSpace, CodeInvalidInput):      errFun(EC50002, IrishubErrMsg)(e.New("validator address is nil")),
	sdkCode(StakeCodeSpace, CodeValidatorJailed):   errFun(EC50002, IrishubErrMsg)(e.New("validator jailed")),

	sdkCode(DistrCodeSpace, CodeInvalidInput):       errFun(EC50002, IrishubErrMsg)(e.New("no delegation distribution info")),
	sdkCode(DistrCodeSpace, CodeNoDistributionInfo): errFun(EC50002, IrishubErrMsg)(e.New("no delegation distribution info")),
}

func errFun(errCode uint32, errMsg string) func(errMsg error) IrisError {
	return func(err error) IrisError {
		msg := errMsg
		if err != nil {
			msg = fmt.Sprintf(errMsg, err.Error())
		}
		return IrisError{
			ErrCode: errCode,
			ErrMsg:  msg,
		}
	}
}
