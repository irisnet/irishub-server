package rpc

import (
	"fmt"
	"github.com/irisnet/irishub-server/errors"
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/rpc/filters"
)

type Filter interface {
	Check(req interface{}) (bool, errors.IrisError)
}

const GlobalFilterPath = "*"

var fts = make(map[string][]Filter, 0)

func RegisterFilters(request string, fs []Filter) {
	if _, ok := fts[request]; ok {
		logger.Error.Println("duplicate registration filter:", request)
	}
	if len(fs) == 0 {
		logger.Error.Println("invalid filter handler", request)
	}
	fts[request] = fs
}

func init() {
	RegisterFilters(GlobalFilterPath, []Filter{filters.HealthCheck{}})
}

func DoFilters(req string) (bool, errors.IrisError) {
	logger.Info.Println(fmt.Sprintf("doFilters req:%s", req))
	//check global filters
	globalFilters, ok := fts[GlobalFilterPath]
	if ok {
		for _, f := range globalFilters {
			ok, err := f.Check(req)
			if !ok {
				return false, err
			}
		}
	}

	//check custom filters
	customFilters, ok := fts[req]
	if ok {
		for _, f := range customFilters {
			ok, err := f.Check(req)
			if !ok {
				return false, err
			}
		}
	}

	return true, errors.IrisError{}
}
