namespace go model

/** coin struct
 * @param amount, token amount
 * @param denom, name of token
 */
struct Coin {
	1: double amount,
	2: string denom
}

/** address structure
 * @param chain, blockchain identify
 * @param app, reserved field
 * @param addr, address in blockchain
 */
struct Address {
	1: string chain,
	2: string app,
	3: string addr
}

/** fee structure
 * @param amount, token amount
 * @param denom, name of token
 */
struct Fee {
	1: double amount,
	2: string denom
}

/** memo structure
 * @param id,
 * @param text, content of memo
 */
struct Memo {
	1: i64 id,
	2: binary text
}

struct GasPrice {
    1: double minGasPrice,
    2: double maxGasPrice,
    3: double avgGasPrice,
    4: string denom
}

/** tx structure
 *
 */
struct Tx {
    1: i64 sequence,
	2: Address sender,
	3: Address receiver,
	4: list<Coin> amount,
	5: Fee fee,
	6: Memo memo,
	7: string type,

	8: string txHash,
	9: string time,
	10: i64 height,
	11: string status,
	12: binary ext,
	13: double gasLimit,
	14: double gasUsed,
	15: Fee actualFee
}

/** common exception
 * @param errCode, error code
 * @param errMsg, error message
 */
exception Exception {
  1: i32 errCode,
  2: string errMsg
}

// ========================================
// define each method request and response
// ========================================

/** txGas request
 * @param txType, txType
 */
struct TxGasRequest {
    1: string txType
}

/** txGas response
 * @param gas, suggest gas
 * @param gasPrice, suggest gasPrice
 */
struct TxGasResponse {
    1: string txType,
    2: double gasLimit,
    3: GasPrice gasPrice
}

/** sequence request
 * @param address, address
 */
struct SequenceRequest {
	1: string address
}

/** sequence response
 * @param sequence, sequence of address
 */
struct SequenceResponse {
	1: i64 sequence
	2: binary ext
}

/** buildTx request
 * @param tx
 */
struct BuildTxRequest {
	1: Tx tx
}

/** buildTx response
 * @param data, result of buildTx
 */
struct BuildTxResponse {
	1: binary data
	2: binary ext
}

/** postTx request
 * @param tx, tx which has been signed
 */
struct PostTxRequest {
	1: binary tx
}

/** postTx response
 * @param txHash, transaction hash
 */
struct PostTxResponse {
	1: string txHash
}

/** SimulateTxRequest request
 * @param tx, tx which has been signed
 */
struct SimulateTxRequest {
	1: binary tx
}

struct RewardDetail {
    1: string valAddr
    2: string name
    3: Coin   amount
}
/** SimulateTxRequest request
 * @param tx, tx which has been signed
 */
struct SimulateTxResponse {
    1: i64  gas
    2: list<RewardDetail> details
}

/** balance request
 * @param address, address of blockchain
 */
struct BalanceRequest {
	1: string address
}

/** balance response
 * @param coins, balance of address
 */
struct BalanceResponse {
	1: list<Coin> coins
}

/** txList request
 * @param address, address of blockchain
 * @param page, current page
 * @param perPage, num of record each page
 * @param status, tx status
 * @param type, tx type
 * @param startTime, tx time
 * @param endTime, tx time
 * @param sort, sort
 * @param q, content of query
 */
struct TxListRequest {
	1: string address,
	2: required i64 page,
	3: required i64 perPage,
	4: string status,
	5: string type,
	6: string startTime,
	7: string endTime,
	8: string sort,
	9: string q,
	10: binary ext,
	11: i64 height
}

/** txDetail request
 * @param txHash, tx hash
 */
struct TxDetailRequest {
	1: required string txHash
}

struct Reward {
    1: string delAddr
    2: list<Coin> amount,
    3: Fee fee,
    4: Memo memo,
    5: string type,
    6: string txHash,
    7: string time,
    8: i64 height,
    9: string status,
    10:string withdrawAddr,
    11: list<RewardDetail> details
}

struct RewardInfoRequest {
	1: string delAddr
	2: string valAddr,
}

struct RewardInfoResponse {
	1: string delAddr
	2: string withdrawAddr
	3: Coin totalRetrieveReward
	4: list<Reward> rewards
}




