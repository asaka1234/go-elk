package go_elk

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-elk/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

// 下单
func (cli *Client) CryDeposit(req ELKCryDepositReq) (*ELKCryDepositRsp, error) {

	rawURL := cli.Params.CryDepositUrl

	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//补充字段
	params["uid"] = cli.Params.MerchantId

	// Generate signature
	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["signature"] = signStr

	var result ELKCryDepositRsp

	resp2, err := cli.ryClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).
		SetCloseConnection(true).
		R().
		SetBody(params).
		SetHeaders(getHeaders()).
		SetDebug(cli.debugMode).
		SetResult(&result).
		SetError(&result).
		Post(rawURL)

	restLog, _ := jsoniter.ConfigCompatibleWithStandardLibrary.Marshal(utils.GetRestyLog(resp2))
	cli.logger.Infof("PSPResty#elk#crydeposit->%s", string(restLog))

	if err != nil {
		return nil, err
	}

	if resp2.StatusCode() != 200 {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("status code: %d", resp2.StatusCode())
	}

	if resp2.Error() != nil {
		//反序列化错误会在此捕捉
		return nil, fmt.Errorf("%v, body:%s", resp2.Error(), resp2.Body())
	}

	return &result, nil
}
