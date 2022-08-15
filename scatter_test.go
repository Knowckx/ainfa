package infa

import (
	"fmt"
	"testing"
	"time"
)

func TestSHTime(t *testing.T) {
	got := SHTime(time.Now().UTC())
	fmt.Println(got)
}
