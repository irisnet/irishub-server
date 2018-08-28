package constants

// environment
const ENV_DEV = "dev"
const ENV_STAGE = "stage"
const ENV_PRO = "pro"

// environment name
const ENV_NAME_ENV = "ENV"
const ENV_NAME_DB_HOST = "DB_HOST"
const ENV_NAME_DB_PORT = "DB_PORT"
const ENV_NAME_DB_User = "DB_USER"
const ENV_NAME_DB_Passwd = "DB_PASSWD"
const ENV_NAME_DB_DATABASE = "DB_DATABASE"
const ENV_NAME_ADDR_NODE_SERVER = "ADDR_NODE_SERVER"

// response status
const STATUS_CODE_OK = 200
const StatusCodeNotContent = 204
const StatusCodeBadRequest = 400

// time layout
const TIME_START = "1970-01-01 00:00:00"
const TIME_LAYOUT_FULL = "2006-01-02 15:04:05"

// define tx type
const TxTypeCoinReceive = "receive"
const TxTypeCoinSend = "send"
const TxTypeStake = "stake"
const TxTypeStakeDelegate = "delegate"
const TxTypeStakeUnbond = "unbond"
const TxTypeStakeBeginUnBonding = "beginUnbonding"
const TxTypeStakeCompleteUnBonding = "completeUnbonding"

// define tx type store in db
const DbTxTypeTransfer = "Transfer"
const DbTxTypeStakeDelegate = "Delegate"
const DbTxTypeStakeBeginUnBonding = "BeginUnbonding"
const DbTxTypeStakeCompleteUnBonding = "CompleteUnbonding"

var TxTypeFrontMapDb = map[string]string{
	TxTypeCoinReceive:            DbTxTypeTransfer,
	TxTypeCoinSend:               DbTxTypeTransfer,
	TxTypeStakeDelegate:          DbTxTypeStakeDelegate,
	TxTypeStakeBeginUnBonding:    DbTxTypeStakeBeginUnBonding,
	TxTypeStakeCompleteUnBonding: DbTxTypeStakeCompleteUnBonding,
}

// define tx status
const TxStatusSuccess = "success"

// define token denom
const Denom = "iris"

// define blockchainRPC uri
const UriBlockChainRPC = "/blockchain"
const UriIrisHubRpc = "/irishub"

// define success status code and fail status code
var SuccessStatusCodes = []int{200}
var ErrorStatusCodes = []int{400, 401, 403, 404, 500}

// define uri of server which expose by block chain
const HttpUriBuildCoinTx = "/build/send"
const HttpUriBuildDelegateTx = "/build/stake/delegate"
const HttpUriBuildUnBondTx = "/build/stake/unbond"
const HttpUriByteTx = "/byteTx"
const HttpUriPostTx = "/tx/send"
const HttpUriGetSequence = "/accounts/%s"             // accounts/{{address}}
const HttpUriGetBalance = "/accounts/%s"              // accounts/{{address}}
const HttpUriGetExRate = "/stake/validator/%s/exRate" // /stake/validator/{address}/exRate
// define http header
const HeaderContentTypeJson = "application/json;charset=utf-8"

// define default tx gas and gasPrice
const DefaultMinGasPrice = 20000000000
const DefaultAvgGasPrice = 30000000000
const DefaultMaxGasPrice = 40000000000
const DefaultTxGasTransfer = 8000
const DefaultTxGasDelegate = 30000
const DefaultTxGasBeginUbonding = 30000
const DefaultTxGasCompleteUnbonding = 30000
