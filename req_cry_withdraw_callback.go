package go_elk

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-elk/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

func (cli *Client) CryWithdrawCallback(req ELKCryWithdrawBackReq, processor func(ELKCryWithdrawBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)

	// Verify signature
	flag, err := utils.Verify(params, cli.Params.BackKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("ELKCry back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
