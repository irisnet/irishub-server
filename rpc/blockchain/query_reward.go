package blockchain

import (
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	commonProtoc "github.com/irisnet/irisnet-rpc/common/codegen/server/model"
	"golang.org/x/net/context"
)

type QueryRewardInfoHandler struct {
}

func (c QueryRewardInfoHandler) Handler(ctx context.Context, request *commonProtoc.RewardInfoRequest) (
	*commonProtoc.RewardInfoResponse, error) {

	reqVO := c.buildRequest(request)
	resVO, err := accountService.GetRewardInfo(reqVO)

	if err.IsNotNull() {
		return nil, BuildException(err)
	}

	return c.buildResponse(resVO, request.ValAddr), nil
}

func (c QueryRewardInfoHandler) buildRequest(req *commonProtoc.RewardInfoRequest) vo.RewardInfoReqVO {
	reqVO := vo.RewardInfoReqVO{
		ValAddr: req.ValAddr,
		DelAddr: req.DelAddr,
	}

	return reqVO
}

func (c QueryRewardInfoHandler) buildResponse(resVO vo.RewardInfoResVo, valAddr string) *commonProtoc.RewardInfoResponse {
	txList := resVO.Txs
	resp := &commonProtoc.RewardInfoResponse{
		DelAddr: resVO.DelAddr,
	}
	if len(txList) == 0 {
		return resp
	}

	var rewards []*commonProtoc.Reward
	var totalRetrieveReward commonProtoc.Coin

	for _, tx := range txList {
		var details []*commonProtoc.RewardDetail
		valAddrMap := make(map[string]string)
		kvPair := simulateTxService.ConvertToTags(tx.Tags)
		records, err := simulateTxService.ParseTags(kvPair)
		amount := BuildCoinsRes(tx.Amount)
		if err == nil && tx.Status == constants.TxStatusSuccess {
			for _, r := range records {
				details = append(details, &commonProtoc.RewardDetail{
					ValAddr: r.ValAddress,
					Name:    r.Name,
					Amount: &commonProtoc.Coin{
						Amount: r.Amount.Amount,
						Denom:  r.Amount.Denom,
					},
				})
				valAddrMap[r.ValAddress] = r.ValAddress
			}

			//如果用户查询指定validator上的提现记录，过滤掉
			if len(valAddr) > 0 {
				_, ok := valAddrMap[valAddr]
				if !ok {
					continue
				}
			}
		}

		if tx.Status == constants.TxStatusSuccess && amount != nil {
			totalRetrieveReward = AddCoin(totalRetrieveReward, *amount[0])
		}

		withdrawAddr := tx.Tags["withdraw-address"]
		reward := &commonProtoc.Reward{
			DelAddr:      tx.From,
			Amount:       amount,
			Fee:          BuildActualFeeRes(tx.ActualFee),
			Memo:         BuildMemoRes(tx.Memo),
			Type:         tx.Type,
			TxHash:       tx.TxHash,
			Time:         tx.Time.String(),
			Height:       tx.Height,
			Status:       tx.Status,
			WithdrawAddr: withdrawAddr,
			Details:      details,
		}
		rewards = append(rewards, reward)
	}

	resp.Rewards = rewards
	resp.TotalRetrieveReward = &totalRetrieveReward

	return resp
}
