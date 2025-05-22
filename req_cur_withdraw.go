package go_elk

import (
	"crypto/tls"
	"encoding/json"
	"github.com/asaka1234/go-elk/utils"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) CurWithdraw(req ELKCurWithdrawReq) (*ELKCurWithdrawRsp, error) {

	rawURL := cli.CurWithdrawUrl

	// 2. Convert struct to map for signing
	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["uid"] = cli.MerchantID //要参与签名计算

	// Generate signature
	signStr, _ := utils.Sign(params, cli.AccessKey)
	params["signature"] = signStr

	// Prepare request
	jsonReq, _ := json.Marshal(req)
	cli.logger.Infof("ELKCurService#withdraw#json: %s", jsonReq)

	// Send HTTP request

	var result ELKCurWithdrawRsp

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
