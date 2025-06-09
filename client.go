package go_elk

import (
	"github.com/asaka1234/go-elk/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	Params *ELKInitParams

	ryClient  *resty.Client
	debugMode bool //是否调试模式
	logger    utils.Logger
}

func NewClient(logger utils.Logger, params *ELKInitParams) *Client {
	return &Client{
		Params: params,

		ryClient:  resty.New(), //client实例
		debugMode: false,
		logger:    logger,
	}
}

func (cli *Client) SetDebugModel(debugModel bool) {
	cli.debugMode = debugModel
}

func (cli *Client) SetMerchantInfo(merchant MerchantInfo) {
	cli.Params.MerchantInfo = merchant
}
