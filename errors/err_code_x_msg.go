package errors

const EC40001 = 40001
const EM40001 = "错误的请求参数: "

const EC40002 = 40002
const EM40002 = "参数转化异常: "

const EC50001 = 50001
const EM50001 = "系统异常: "

const EC60001 = 60001
const EM60001 = "交易已存在 "

const EC60002 = 60002
const EM60002 = "交易超时"

const EC60003 = 60003
const EM60003 = "gas不足"

var (
	InvalidReqParamsErr = IrisErr(EC40001, EM40001)
	ParamConvertErr     = IrisErr(EC40002, EM40002)
	SysErr              = IrisErr(EC50001, EM50001)
	TxExistedErr        = IrisErr(EC60001, EM60001)
	TxTimeoutErr        = IrisErr(EC60002, EM60002)
	OutOfGasErr         = IrisErr(EC60003, EM60003)
)

//sdkCode
const sdkCodeOutOfGas = uint16(12)

var sdkCodeToIrisCodeMap = map[uint16]IrisError{
	sdkCodeOutOfGas: OutOfGasErr(nil),
}
