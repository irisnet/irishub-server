package services

import (
	"bytes"
	"encoding/json"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/vo"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/helper"
	"strings"
)

type SimulateTxService struct {
}

type kvPair struct {
	TagKey   string `json:"tag_key"`
	TagValue string `json:"tag_value"`
}
type abciResult struct {
	Code      uint32   `json:"code"`
	Data      []byte   `json:"data"`
	Log       string   `json:"log"`
	GasWanted string   `json:"gas_wanted"`
	GasUsed   string   `json:"gas_used"`
	FeeAmount string   `json:"fee_amount"`
	FeeDenom  string   `json:"fee_denom"`
	Tags      []kvPair `json:"tags"`
}
type simulateResult struct {
	GasEstimate string     `json:"gas_estimate"`
	Result      abciResult `json:"result"`
}

func (s SimulateTxService) SimulateTx(reqVO vo.SimulateTxReqVO) (vo.SimulateTxResVO, errors.IrisError) {
	var resVO vo.SimulateTxResVO

	tx := reqVO.Tx

	reqPostTx := bytes.NewBuffer(tx)

	resVo, err := s.simulate(reqPostTx)
	if err.IsNotNull() {
		return resVO, err
	}

	return resVo, irisErr
}

func (s SimulateTxService) simulate(requestBody *bytes.Buffer) (res vo.SimulateTxResVO, irisErr errors.IrisError) {
	resByte, err := broadcastTx(false, true, requestBody)
	if err.IsNotNull() {
		return res, err
	}

	var resp simulateResult

	er := json.Unmarshal(resByte, &resp)
	if er != nil {
		return res, errors.SysErr(er)
	}

	if resp.Result.Code != 0 {
		logger.Error.Printf("%v: err is %v\n", "simulate", resp.Result.Log)
		return res, NewIrisErr(resp.Result.Code, resp.Result.Log, nil)
	}

	records, er := s.ParseTags(resp.Result.Tags)
	if er != nil {
		return res, errors.SysErr(er)
	}
	res.Gas = helper.ConvertStrToInt(resp.GasEstimate)
	res.Records = records

	return res, irisErr
}

func (s SimulateTxService) ParseTags(tags []kvPair) (records []vo.Record, err error) {
	var txType string
	for _, tag := range tags {
		if tag.TagKey == constants.TagNmAction {
			txType = tag.TagValue
			break
		}
	}

	switch txType {
	case constants.TagNmDistributionWithdrawDelegatorRewardsAll:
		var valAddr string
		var amt *vo.Coin
		for _, tag := range tags {
			if strings.HasPrefix(tag.TagKey, constants.TagNmDistributionWithdrawRewardFromValidator) {
				valAddr = strings.TrimPrefix(tag.TagKey, constants.TagNmDistributionWithdrawRewardFromValidator)
				denom, amount, err := helper.ParseCoin(tag.TagValue)
				if err != nil {
					return nil, err
				}
				amt = &vo.Coin{
					Denom:  denom,
					Amount: helper.ConvertStrToFloat(amount),
				}
				validator, err := candidateModel.GetCandidateDetail(valAddr)
				if err != nil {
					return nil, err
				}
				records = append(records, vo.Record{
					ValAddress: valAddr,
					Name:       validator.Description.Moniker,
					Amount:     amt,
				})
			}
		}
		return records, nil
	case constants.TagNmDistributionWithdrawDelegationReward:
		var valAddr string
		var amt *vo.Coin
		for _, tag := range tags {
			if tag.TagKey == constants.TagNmDistributionSourceValidator {
				valAddr = tag.TagValue
			} else if tag.TagKey == constants.TagNmDistributionWithdrawRewardTotal {
				denom, amount, err := helper.ParseCoin(tag.TagValue)
				if err != nil {
					return nil, err
				}
				amt = &vo.Coin{
					Denom:  denom,
					Amount: helper.ConvertStrToFloat(amount),
				}
			}
		}
		validator, err := candidateModel.GetCandidateDetail(valAddr)
		if err != nil {
			return nil, err
		}
		records = append(records, vo.Record{
			ValAddress: valAddr,
			Name:       validator.Description.Moniker,
			Amount:     amt,
		})
	}
	return records, nil
}

func (s SimulateTxService) ConvertToTags(tagMap map[string]string) (kv []kvPair) {
	if len(tagMap) == 0 {
		return
	}
	for k, v := range tagMap {
		kv = append(kv, kvPair{
			TagKey:   k,
			TagValue: v,
		})
	}
	return kv
}
