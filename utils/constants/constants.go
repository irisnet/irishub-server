package constants

// environment
const ENV_DEV = "dev"
const ENV_STAGE  = "stage"
const ENV_PRO = "pro"

// environment name
const ENV_NAME_ENV  = "ENV"
const ENV_NAME_DB_HOST  = "DB_HOST"
const ENV_NAME_DB_PORT  = "DB_PORT"
const ENV_NAME_ADDR_NODE_SERVER = "ADDR_NODE_SERVER"

const PAGE_LIMIT_NUM = 20

// response status
const STATUS_CODE_OK = 200
const STATUS_SUCCESS = "success"
const STATUS_FAIL = "fail"

// time layout
const TIME_START  = "1970-01-01 00:00:00"
const TIME_LAYOUT_FULL = "2006-01-02 15:04:05"

// define tx type
const TxTypeCoinReceive  = "receive"
const TxTypeCoinSend  = "send"
const TxTypeStakeDelegate  = "delegate"
const TxTypeStakeUnBond  = "unbond"

// define tx type store in db
const DbTxTypeCoin = "coin"
const DbTxTypeStakeDelegate  = "delegate"
const DbTxTypeStakeUnBond = "unbond"
var TxTypeFrontMapDb = map[string]string{
	TxTypeCoinReceive: DbTxTypeCoin,
	TxTypeCoinSend: DbTxTypeCoin,
	TxTypeStakeDelegate: DbTxTypeStakeDelegate,
	TxTypeStakeUnBond: DbTxTypeStakeUnBond,
}

// define token denom
const Denom  = "iris"

// define success status code and fail status code
var SuccessStatusCodes  = []int{200}
var ErrorStatusCodes = []int{400}


// define uri of server which expose by block chain
const HttpUriBuildTx = "/build/send"
const HttpUriByteTx  = "/byteTx"
const HttpUriPostTx  = "/tx"

// define http header
const HeaderContentTypeJson = "application/json;charset=utf-8"