package web

import (
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/tidwall/gjson"
	"gopkg.in/resty.v1"
)

type RestClient struct {
	*resty.Request
}

func NewRestClient() *RestClient {
	out := &RestClient{}
	client := resty.New().SetTimeout(10 * time.Second)
	req := client.R().SetHeader("Accept", "application/json")
	out.Request = req
	return out
}

func (cli *RestClient) DoGet(url string, queryParams map[string]string) (*gjson.Result, error) {
	log.Debug().Str("url,", url).Msg("send request get")

	if queryParams != nil {
		cli.Request = cli.SetQueryParams(queryParams)
	}
	resp, err := cli.Get(url)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	err = CheckRestyResponse(resp)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	res := gjson.ParseBytes(resp.Body())
	return &res, nil
}

func CheckRestyResponse(resp *resty.Response) error {
	if resp.IsError() {
		log.Error().Int("Status Code", resp.StatusCode()).Msg("Http Request Failed")
		data := resp.Body()
		log.Error().Str("Body", string(data)).Send()
		return fmt.Errorf("Http Request Failed")
	}
	log.Debug().Int("Status Code", resp.StatusCode()).Msg("Http Request Success.")
	return nil
}
