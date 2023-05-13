package format

import (
	"fmt"
	"testing"

	"github.com/Knowckx/infa/path"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
)

func Test_SaveJson(t *testing.T) {
	testMap := map[string]string{
		"AA":  "BB",
		"AA1": "BB1",
	}
	err := SaveYamlFile("asu", testMap)
	assert.Nil(t, err)
}

/*
Yaml解析的字段匹配需要注意下，结构体字段我们习惯大写，
但是对应的Yaml文件的key需要是小写这个包才会自动匹配。
假如两边都是大写，注释`yaml:"URL"`就是必需的
嵌套的结构体：`yaml:",inline"`
*/
type ClientInfo struct {
	URL          string `yaml:"URL"`
	ClientID     string `yaml:"ClientID"`
	ClientSecret string `yaml:"ClientSecret"`
}

func Test_LoadYamlFile(t *testing.T) {
	out := &ClientInfo{}
	fPath := path.GetInfaPath("testfiles", "ClientInfo.yaml")
	err := LoadYamlFile(fPath, out)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v", out)
	return
}

func Test_SaveYamlFile(t *testing.T) {
	info := &ClientInfo{}
	info.URL = "1.2"
	info.ClientID = "1"
	info.ClientSecret = "3"
	bys, err := yaml.Marshal(info)
	if err != nil {
		log.Error().Err(err).Msg("yaml.Marshal error")
	}
	fmt.Println("final\n", string(bys))
	return
}
