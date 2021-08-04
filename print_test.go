package infa

import (
	"testing"
)

func TestPrint(t *testing.T) {
	au1 := AuthParams{}
	au1.Host = "host1"
	au1.User = "user2"
	au2 := AuthParams{}
	au2.Host = "host1"
	au2.User = "user2"

	in := []*AuthParams{
		&au1, &au2,
	}
	// Printf("%v", in)
	// Print(in)
	PrintToml(in)
}

func Test_Printf(t *testing.T) {
	Printf("123 %s", "qq")
	Printf("123 %s", "qq")
	Printf("123 %s", "qq")
}
