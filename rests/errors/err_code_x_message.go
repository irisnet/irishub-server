package errors

const EC_BAD_REQUEST = "bad request"
const EM_BAD_REQUEST = "错误的请求参数"

const EC_UNAUTHERIZATION = "unautherization"
const EM_UNAUTHERIZATION = "未认证"

const EC_SYSERR = "system error"

var EC_X_EM map[string]string

func init() {
	EC_X_EM[EC_BAD_REQUEST] = EM_BAD_REQUEST
	EC_X_EM[EC_UNAUTHERIZATION] = EM_UNAUTHERIZATION
}
