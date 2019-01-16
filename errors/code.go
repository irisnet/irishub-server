package errors

import (
	"fmt"
)

const (
	EC10000 = 10000

	EC40000 = 40000
	EC40001 = 40001
	EC40002 = 40002

	EC50000 = 50000
	EC50001 = 50001
	EC50002 = 50002
	EC50003 = 50003
	EC50004 = 50004
	EC50005 = 50005
	EC50006 = 50006
	EC50007 = 50007
	EC50008 = 50008
	EC50009 = 50009
	EC50010 = 50010
	EC50011 = 50011
	EC50012 = 50012
	EC50013 = 50013
	EC50014 = 50014
	EC50015 = 50015
	EC50016 = 50016
	EC50017 = 50017
	EC50018 = 50018
	EC50019 = 50019
	EC50020 = 50020
	EC59999 = 59999
)

var (
	irisErr = map[int]string{
		EC10000: "the system is under maintenance",
		EC40000: "system error: %s",
		EC40001: "invalid param error: %s",
		EC40002: "param convert error: %s",
		EC50000: "external system error: %s",
		EC59999: "tx timeout: %s",
	}

	SysMaintenance   = errFun(EC10000, irisErr[EC10000])
	SysErr           = errFun(EC40000, irisErr[EC40000])
	InvalidParamsErr = errFun(EC40001, irisErr[EC40001])
	ParamConvertErr  = errFun(EC40002, irisErr[EC40002])
	ExtSysUnKnownErr = errFun(EC50000, irisErr[EC50000])
	TimeoutErr       = errFun(EC59999, irisErr[EC59999])
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
	CodeInternal:          errFun(EC50001, irisErr[EC50000]),
	CodeTxDecode:          errFun(EC50002, irisErr[EC50000]),
	CodeInvalidSequence:   errFun(EC50003, irisErr[EC50000]),
	CodeUnauthorized:      errFun(EC50004, irisErr[EC50000]),
	CodeInsufficientFunds: errFun(EC50005, irisErr[EC50000]),
	CodeUnknownRequest:    errFun(EC50006, irisErr[EC50000]),
	CodeInvalidAddress:    errFun(EC50007, irisErr[EC50000]),
	CodeInvalidPubKey:     errFun(EC50008, irisErr[EC50000]),
	CodeUnknownAddress:    errFun(EC50009, irisErr[EC50000]),
	CodeInsufficientCoins: errFun(EC50010, irisErr[EC50000]),
	CodeInvalidCoins:      errFun(EC50011, irisErr[EC50000]),
	CodeOutOfGas:          errFun(EC50012, irisErr[EC50000]),
	CodeMemoTooLarge:      errFun(EC50013, irisErr[EC50000]),
	CodeInsufficientFee:   errFun(EC50014, irisErr[EC50000]),
	CodeOutOfService:      errFun(EC50015, irisErr[EC50000]),
	CodeTooManySignatures: errFun(EC50016, irisErr[EC50000]),
	CodeGasPriceTooLow:    errFun(EC50017, irisErr[EC50000]),
	CodeInvalidGas:        errFun(EC50018, irisErr[EC50000]),
	CodeInvalidTxFee:      errFun(EC50019, irisErr[EC50000]),
	CodeInvalidFeeDenom:   errFun(EC50020, irisErr[EC50000]),
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
