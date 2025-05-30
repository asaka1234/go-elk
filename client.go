package go_elk

import (
	"github.com/asaka1234/go-elk/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params ELKInitParams

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, params ELKInitParams) *Client {
	return &Client{
		Params: params,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchantInfo(merchantId int, accessKey, backKey string) {
	cli.Params.MerchantId = merchantId
	cli.Params.AccessKey = accessKey
	cli.Params.BackKey = backKey
}
