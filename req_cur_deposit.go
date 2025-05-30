package go_elk

import (
	"crypto/tls"
	"encoding/json"
	"github.com/asaka1234/go-elk/utils"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) CurDeposit(req ELKCurDepositReq) (*ELKCurDepositRsp, error) {

	rawURL := cli.Params.CurDepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["uid"] = cli.Params.MerchantId

	// Log request
	cli.logger.Infof("ELKCurService#req: %+v", req)

	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["signature"] = signStr

	// Convert to JSON
	jsonStr, _ := json.Marshal(params)
	cli.logger.Infof("ELKCurService#deposit#json: %s", jsonStr)

	var result ELKCurDepositRsp

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
