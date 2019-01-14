package filters

import (
	"encoding/json"
	"fmt"
	"github.com/irisnet/irishub-server/configs"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
	"github.com/irisnet/irishub-server/utils/http"
)

type health struct {
	Status  string `json:"status"`
	ErrCode string `json:"errCode"`
	ErrMsg  string `json:"errMsg"`
	Data    struct {
		IsMaintained bool `json:"isMaintained"`
	} `json:"data"`
}

type HealthCheck struct{}

func (HealthCheck) Check(req interface{}) (bool, errors.IrisError) {
	reqUrl := fmt.Sprintf("%s%s", configs.ServerConfig.RainbowServer, constants.HttpApiHealthCheck)
	resp := http.Get(reqUrl, nil)
	if resp.Error != nil || resp.Code != constants.STATUS_CODE_OK {
		logger.Error.Println("request api error ", reqUrl)
	}

	var h health
	if err := json.Unmarshal(resp.Data, &h); err != nil {
		logger.Error.Println("healthCheck Unmarshal error")
	}

	if h.Status == "success" && !h.Data.IsMaintained {
		return true, errors.IrisError{}
	}
	return false, errors.SysMaintenance("HealthCheck Filter")
}
