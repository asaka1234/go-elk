package go_elk

import (
	"github.com/asaka1234/go-elk/utils"
	"github.com/go-resty/resty/v2"
)

type Client struct {
	MerchantID int    // merchantId
	AccessKey  string // accessKey
	BackKey    string //backKey

	CurDepositUrl  string
	CurWithdrawUrl string

	CryDepositUrl  string
	CryWithdrawUrl string

	ryClient *resty.Client
	logger   utils.Logger
}

func NewClient(logger utils.Logger, merchantID int, accessKey, backKey, curDepositUrl, curWithdrawUrl, cryDepositUrl, cryWithdrawUrl string) *Client {
	return &Client{
		MerchantID:     merchantID,
		AccessKey:      accessKey,
		BackKey:        backKey,
		CurDepositUrl:  curDepositUrl,
		CurWithdrawUrl: curWithdrawUrl,
		CryDepositUrl:  cryDepositUrl,
		CryWithdrawUrl: cryWithdrawUrl,

		ryClient: resty.New(), //client实例
		logger:   logger,
	}
}

func (cli *Client) SetMerchantId(merchantId int) {
	cli.MerchantID = merchantId
}
