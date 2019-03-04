package common

import (
	"fmt"
	"testing"
)

func TestRDM(t *testing.T) {
	fmt.Println(RandomString(6))
	fmt.Println(RandomNumber(6))
}
