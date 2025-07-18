package go_elk

import (
	"crypto/tls"
	"fmt"
	"github.com/asaka1234/go-elk/utils"
	jsoniter "github.com/json-iterator/go"
	"github.com/mitchellh/mapstructure"
)

func (cli *Client) CryWithdraw(req ELKCryWithdrawReq) (*ELKCryWithdrawRsp, error) {

	rawURL := cli.Params.CryWithdrawUrl

	// 2. Convert struct to map for signing
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	//补充字段
	params["uid"] = cli.Params.MerchantId

	//计算签名
	signStr, _ := utils.Sign(params, cli.Params.AccessKey)
	params["signature"] = signStr

	var result ELKCryWithdrawRsp

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
	cli.logger.Infof("PSPResty#elk#crywithdraw->%+v", string(restLog))

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
