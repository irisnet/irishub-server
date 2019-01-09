package errors

import (
	"errors"
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	fmt.Println(InvalidParamsErr(errors.New("test")))
	fmt.Println(SdkCodeToIrisErr("sdk", uint16(3)))
}
