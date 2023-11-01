package convert

import (
	"fmt"
	"testing"
)

func TestPStringToString(t *testing.T) {
	var pstring *string
	st := "dfefef"
	pstring = &st
	var s string
	s = PStringToString(pstring)
	fmt.Println(s)
}
