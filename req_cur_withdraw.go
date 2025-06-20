package go_elk

import (
	"crypto/tls"
	"github.com/asaka1234/go-elk/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) CurWithdraw(req ELKCurWithdrawReq) (*ELKCurWithdrawRsp, error) {

	rawURL := cli.Params.CurWithdrawUrl

	// 2. Convert struct to map for signing
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//补充字段
	params["uid"] = cli.Params.MerchantId

	//计算签名
	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["signature"] = signStr

	var result ELKCurWithdrawRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
