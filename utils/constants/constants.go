package constants

// environment
const ENV_DEV = "dev"
const ENV_STAGE = "stage"
const ENV_PRO = "pro"

// environment name
const ENV_NAME_ENV = "ENV"
const ENV_NAME_DB_ADDR = "DB_ADDR"
const ENV_NAME_DB_User = "DB_USER"
const ENV_NAME_DB_Passwd = "DB_PASSWD"
const ENV_NAME_DB_DATABASE = "DB_DATABASE"
const ENV_NAME_LCD_SERVER = "LCD_SERVER"
const ENV_NAME_RAINBOW_SERVER = "RAINBOW_SERVER"

// response status
const STATUS_CODE_OK = 200
const StatusCodeNotContent = 204
const StatusCodeBadRequest = 400
const StatusInternalServerError = 500

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
const TxTypeStakeBeginRedelegate = "redelegate"
const TxTypeSetWithdrawAddress = "setWithdrawAddress"
const TxTypeWithdrawDelegatorReward = "withdrawDelegatorReward"
const TxTypeWithdrawDelegatorRewardsAll = "withdrawDelegatorRewardsAll"
const TxTypeWithdrawValidatorRewardsAll = "withdrawValidatorRewardsAll"

const TagNmAction = "action"
const TagNmDistributionWithdrawDelegatorRewardsAll = "withdraw_delegation_rewards_all"
const TagNmDistributionWithdrawDelegationReward = "withdraw_delegation_reward"
const TagNmDistributionWithdrawRewardFromValidator = "withdraw-reward-from-validator-"
const TagNmDistributionSourceValidator = "source-validator"
const TagNmDistributionWithdrawRewardTotal = "withdraw-reward-total"

// define tx type store in db
const DbTxTypeTransfer = "Transfer"
const DbTxTypeStakeDelegate = "Delegate"
const DbTxTypeStakeBeginUnBonding = "BeginUnbonding"
const DbTxTypeBeginRedelegate = "BeginRedelegate"
const DbTxTypeSetWithdrawAddress = "SetWithdrawAddress"
const DbTxTypeWithdrawDelegatorReward = "WithdrawDelegatorReward"
const DbTxTypeWithdrawDelegatorRewardsAll = "WithdrawDelegatorRewardsAll"
const DbTxTypeWithdrawValidatorRewardsAll = "WithdrawValidatorRewardsAll"

var TxTypeFrontMapDb = map[string]string{
	TxTypeCoinReceive:                 DbTxTypeTransfer,
	TxTypeCoinSend:                    DbTxTypeTransfer,
	TxTypeStakeDelegate:               DbTxTypeStakeDelegate,
	TxTypeStakeBeginUnBonding:         DbTxTypeStakeBeginUnBonding,
	TxTypeStakeBeginRedelegate:        DbTxTypeBeginRedelegate,
	TxTypeSetWithdrawAddress:          DbTxTypeSetWithdrawAddress,
	TxTypeWithdrawDelegatorReward:     DbTxTypeWithdrawDelegatorReward,
	TxTypeWithdrawDelegatorRewardsAll: DbTxTypeWithdrawDelegatorRewardsAll,
	TxTypeWithdrawValidatorRewardsAll: DbTxTypeWithdrawValidatorRewardsAll,
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
var ErrorStatusCodes = []int{400, 401, 403, 404}

const HttpUriBroadcastTx = "/tx/broadcast?commit=false&async=%v&simulate=%v"

const HttpUriGetSequence = "/auth/accounts/%s" // auth/{{address}}
const HttpUriGetBalance = "/auth/accounts/%s"  // auth/{{address}}
const HttpUriGetValidators = "/stake/validators/%s"
const HttpUriGetWithdrawAddr = "/distribution/%s/withdrawAddress"
const HttpApiHealthCheck = "/ops_ctl/latest"

// define http header
const HeaderContentTypeJson = "application/json;charset=utf-8"

// define default tx gas and gasPrice
const DefaultMinGasPrice = 20000000000
const DefaultAvgGasPrice = 30000000000
const DefaultMaxGasPrice = 40000000000

const DefaultTxGasTransfer = 200000
const DefaultTxGasDelegate = 200000
const DefaultTxGasBeginUbonding = 200000
const DefaultTxGasCompleteUnbonding = 200000
