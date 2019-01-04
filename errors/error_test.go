package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	fmt.Println(InvalidParamsErr(errors.New("test")))
}
