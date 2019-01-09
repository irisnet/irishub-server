package errors

import (
	"fmt"
)

const (
	EC40000 = 40000
	EC40001 = 40001
	EC40002 = 40002

	EC50000 = 50000
	EC50001 = 50001

	EC60000 = 60000
)

var (
	irisErr = map[int]string{
		EC40000: "system error: %s",
		EC40001: "invalid param error: %s",
		EC40002: "param convert error: %s",
		EC50000: "external system error: %s",
		EC50001: "external system unknown error: %s",
		EC60000: "tx timeout: %s",
	}

	SysErr           = errFun(EC40000, irisErr[EC40000])
	InvalidParamsErr = errFun(EC40001, irisErr[EC40001])
	ParamConvertErr  = errFun(EC40002, irisErr[EC40002])
	ExtSysUnKnownErr = errFun(EC50001, irisErr[EC50001])
	TimeoutErr       = errFun(EC60000, irisErr[EC60000])
)

type CodeType uint32

//sdkCode
const (
	//sdk common code
	CodeInternal          CodeType = 1
	CodeTxDecode          CodeType = 2
	CodeInvalidSequence   CodeType = 3
	CodeUnauthorized      CodeType = 4
	CodeInsufficientFunds CodeType = 5
	CodeUnknownRequest    CodeType = 6
	CodeInvalidAddress    CodeType = 7
	CodeInvalidPubKey     CodeType = 8
	CodeUnknownAddress    CodeType = 9
	CodeInsufficientCoins CodeType = 10
	CodeInvalidCoins      CodeType = 11
	CodeOutOfGas          CodeType = 12
	CodeMemoTooLarge      CodeType = 13
	CodeInsufficientFee   CodeType = 14
	CodeOutOfService      CodeType = 15
	CodeTooManySignatures CodeType = 16
	CodeGasPriceTooLow    CodeType = 17
	CodeInvalidGas        CodeType = 18
	CodeInvalidTxFee      CodeType = 19
	CodeInvalidFeeDenom   CodeType = 20
)

type ErrFunc func(msg string, args ...interface{}) IrisError

var sdkCodeToErrFunc = map[CodeType]ErrFunc{
	CodeInternal:          errFun(EC50000, irisErr[EC50000]),
	CodeTxDecode:          errFun(EC50000, irisErr[EC50000]),
	CodeInvalidSequence:   errFun(EC50000, irisErr[EC50000]),
	CodeUnauthorized:      errFun(EC50000, irisErr[EC50000]),
	CodeInsufficientFunds: errFun(EC50000, irisErr[EC50000]),
	CodeUnknownRequest:    errFun(EC50000, irisErr[EC50000]),
	CodeInvalidAddress:    errFun(EC50000, irisErr[EC50000]),
	CodeInvalidPubKey:     errFun(EC50000, irisErr[EC50000]),
	CodeUnknownAddress:    errFun(EC50000, irisErr[EC50000]),
	CodeInsufficientCoins: errFun(EC50000, irisErr[EC50000]),
	CodeInvalidCoins:      errFun(EC50000, irisErr[EC50000]),
	CodeOutOfGas:          errFun(EC50000, irisErr[EC50000]),
	CodeMemoTooLarge:      errFun(EC50000, irisErr[EC50000]),
	CodeInsufficientFee:   errFun(EC50000, irisErr[EC50000]),
	CodeOutOfService:      errFun(EC50000, irisErr[EC50000]),
	CodeTooManySignatures: errFun(EC50000, irisErr[EC50000]),
	CodeGasPriceTooLow:    errFun(EC50000, irisErr[EC50000]),
	CodeInvalidGas:        errFun(EC50000, irisErr[EC50000]),
	CodeInvalidTxFee:      errFun(EC50000, irisErr[EC50000]),
	CodeInvalidFeeDenom:   errFun(EC50000, irisErr[EC50000]),
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
