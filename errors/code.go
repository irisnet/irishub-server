package errors

import (
	"fmt"
)

const (
	IrisHubErrMsg = "irishub error:%s"

	EC40001 = 40001
	EC40002 = 40002

	EC50001 = 50001
	EC50002 = 50002

	EC60001 = 60001
	EC60002 = 60002
	EC60003 = 60003
	EC60004 = 60004
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

type ErrFunc func(msg string, args ...interface{}) IrisError

var sdkCodeToErrFunc = map[string]ErrFunc{
	sdkCode(RootCodeSpace, CodeInternal):        errFun(EC50002, IrisHubErrMsg),
	sdkCode(RootCodeSpace, CodeUnauthorized):    errFun(EC50002, IrisHubErrMsg),
	sdkCode(RootCodeSpace, CodeUnknownRequest):  errFun(EC50002, IrisHubErrMsg),
	sdkCode(RootCodeSpace, CodeInvalidAddress):  errFun(EC50002, IrisHubErrMsg),
	sdkCode(RootCodeSpace, CodeInvalidSequence): errFun(EC50002, IrisHubErrMsg),
	//distinct
	sdkCode(RootCodeSpace, CodeOutOfGas): errFun(EC60003, IrisHubErrMsg),

	sdkCode(StakeCodeSpace, CodeInvalidValidator):  errFun(EC50002, IrisHubErrMsg),
	sdkCode(StakeCodeSpace, CodeInvalidDelegation): errFun(EC60004, IrisHubErrMsg),
	sdkCode(StakeCodeSpace, CodeInvalidInput):      errFun(EC50002, IrisHubErrMsg),
	sdkCode(StakeCodeSpace, CodeValidatorJailed):   errFun(EC50002, IrisHubErrMsg),

	sdkCode(DistrCodeSpace, CodeInvalidInput):       errFun(EC50002, IrisHubErrMsg),
	sdkCode(DistrCodeSpace, CodeNoDistributionInfo): errFun(EC50002, IrisHubErrMsg),
}

func errFun(errCode uint32, errMsg string) func(msg string, args ...interface{}) IrisError {
	return func(msg string, args ...interface{}) IrisError {
		message := msg
		if len(args) > 0 {
			message = fmt.Sprintf(msg, args)
		}
		msg = fmt.Sprintf(errMsg, message)
		return IrisError{
			ErrCode: errCode,
			ErrMsg:  msg,
		}
	}
}
