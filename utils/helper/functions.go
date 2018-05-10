package helper

import (
	"encoding/json"

	"github.com/irisnet/iris-api-server/modules/logger"
)

func ToJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		logger.Error.Println(err)
	}
	return string(data)
}
