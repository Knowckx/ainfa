package infa

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Printf(t *testing.T) {
	Printf("123 %s", "qq")
	Printf("123 %s", "qq")
	Printf("123 %s", "qq")
}

func Test_RecordExecTime(t *testing.T) {
	defer RecordExecTime("Test_RecordExecTime")()
	time.Sleep(2 * time.Second)
}

type User struct {
	Name string
	Age  int
	Home *Home
}

type Home struct {
	Address     string
	PeopleLives []string
	PeopleInfo  map[string]string
}

// 输出一个结构体嵌套的变量
func GenTestUser() *User {
	ho := Home{}
	ho.Address = "beijing"
	ho.PeopleLives = []string{
		"Raman Kalita", "Abhishek Garg",
	}
	ho.PeopleInfo = map[string]string{
		"Raman Kalita":  "civil servant",
		"Abhishek Garg": "teacher",
	}
	us := User{}
	us.Name = "Alise"
	us.Age = 31
	us.Home = &ho
	return &us
}

func Test_PrintJson(t *testing.T) {
	us := GenTestUser()
	PrintJson(us)
}

func Test_Func1(t *testing.T) {
	out, err := Func1()
	if err != nil {
		fmt.Printf("%+v", err)
	}
	assert.Nil(t, err)
	fmt.Println(out)
}
func Func1() (int, error) {
	s1 := make([]int, 3)
	s2 := make([]int, 0, 3)
	fmt.Println(s1, s2)
	return 0, nil
}
