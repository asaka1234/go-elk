package go_elk

import (
	"crypto/tls"
	"github.com/asaka1234/go-elk/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) CryWithdraw(req ELKCryWithdrawReq) (*ELKCryWithdrawRsp, error) {

	rawURL := cli.CryWithdrawUrl

	// 2. Convert struct to map for signing
	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["uid"] = cli.MerchantID //uid要参与签名

	// 3. Generate signature
	signStr, _ := utils.Sign(params, cli.AccessKey)
	params["signature"] = signStr

	var result ELKCryWithdrawRsp

	_, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	//fmt.Printf("result: %s\n", string(resp.Body()))

	if err != nil {
		return nil, err
	}

	return &result, nil
}
