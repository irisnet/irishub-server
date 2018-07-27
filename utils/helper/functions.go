package helper

import (
	"encoding/json"
	"strings"
	"time"
	
	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
	"fmt"
	"strconv"
	"math"
)

// convert object to json
func ToJson(v interface{}) string {
	data, err := json.Marshal(v)
	if err != nil {
		logger.Error.Println(err)
	}
	return string(data)
}

// parse str to time
func ParseTime(layout string, str string) (time.Time, error)  {
	return time.Parse(layout, str)
}

// parse str to time and layout is "2006-01-02 15:04:05"
func ParseFullTime(str string) (time.Time, error) {
	return time.Parse(constants.TIME_LAYOUT_FULL, str)
}

// parse query param: sort
// convert string to slice
func ParseParamSort(sort string) []string {
	var (
		sorts []string
	)
	if sort != "" {
		sorts = strings.Split(sort, ",")
	}
	
	return sorts
}

// parse page param
// get skip, limit variable which used in database
func ParseParamPage(page int, perPage int) (skip int, limit int)  {
	if page == 0 && perPage == 0 {
		perPage = 10
	}
	return (page - 1) * perPage, perPage
}

// contains method for a slice
func SliceContains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func ConvertRatStrToFloat(rat string) (float64, error)  {
	var (
		numerator, denominator float64
		err error
	)
	rats := strings.Split(rat, "/")
	if len(rats) != 2 {
		return float64(0), fmt.Errorf("invalid rat string: %v", rat)
	}
	numerator, err = strconv.ParseFloat(rats[0], 64)
	if err != nil {
		return float64(0), err
	}
	denominator, err = strconv.ParseFloat(rats[1], 64)
	if err != nil {
		return float64(0), err
	}
	return numerator/denominator, nil
}

func ConvertFloatToInt(f float64) int64 {
	return int64(math.Floor(f + 0.5))
}

