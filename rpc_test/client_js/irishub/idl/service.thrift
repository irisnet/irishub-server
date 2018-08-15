include "model.thrift"

namespace go model

service IRISHubService {

	list<model.Candidate> GetCandidateList(1: model.CandidateListRequest req) throws (1:model.Exception e),

	model.Candidate GetCandidateDetail(1: model.CandidateDetailRequest req) throws (1:model.Exception e),

	model.ValidatorExRateResponse GetValidatorExRate(1: model.ValidatorExRateRequest req) throws (1:model.Exception e),

	list<model.Candidate> GetDelegatorCandidateList(1: model.DelegatorCandidateListRequest req) throws (1:model.Exception e),

	model.TotalShareResponse GetDelegatorTotalShares(1: model.TotalShareRequest req) throws (1:model.Exception e),
}