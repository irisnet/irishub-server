package errors

import (
	"fmt"
	"testing"
)

func TestErrors(t *testing.T) {
	fmt.Println(InvalidParamsErr("test"))
	fmt.Println(SdkCodeToIrisErr("sdk", uint16(3), ""))
}
