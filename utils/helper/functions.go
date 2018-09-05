package helper

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/irisnet/irishub-server/modules/logger"
	"github.com/irisnet/irishub-server/utils/constants"
	"strconv"
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
func ParseTime(layout string, str string) (time.Time, error) {
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
func ParseParamPage(page int, perPage int) (skip int, limit int) {
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

func ConvertStrToInt(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		logger.Error.Println("Convert str to int failed")
	}
	return i
}

func ConvertStrToFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		logger.Error.Println("Convert str to int failed")
	}
	return i
}
