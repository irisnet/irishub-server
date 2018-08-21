include "model.thrift"

namespace go model

service BlockChainService {
	// get tx gas
	model.TxGasResponse GetTxGas(1: model.TxGasRequest req) throws (1:model.Exception e),

    // get sequence
	model.SequenceResponse GetSequence(1: model.SequenceRequest req) throws (1:model.Exception e),

	// build tx
	model.BuildTxResponse BuildTx(1: model.BuildTxRequest req) throws (1:model.Exception e),

	// post tx
	model.PostTxResponse PostTx(1: model.PostTxRequest req) throws (1:model.Exception e),

	// get balance
	model.BalanceResponse GetBalance(1: model.BalanceRequest req) throws (1:model.Exception e),

	// get tx list
	list<model.Tx> GetTxList(1: model.TxListRequest req) throws (1:model.Exception e),

	// get tx detail
	model.Tx GetTxDetail(1: model.TxDetailRequest req) throws (1:model.Exception e),
}
