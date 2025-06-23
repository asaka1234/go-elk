package go_elk

import (
	"encoding/json"
	"errors"
	"github.com/asaka1234/go-elk/utils"
	"github.com/mitchellh/mapstructure"
	"log"
)

// crypto的充值回调
func (cli *Client) CryDepositCallback(req ELKCryDepositBackReq, sign string, processor func(ELKCryDepositBackReq) error) error {
	//验证签名
	var params map[string]interface{}
	mapstructure.Decode(req, &params)
	params["signature"] = sign //保存一下从header传过来的原始签名sign

	// Verify signature
	flag, err := utils.Verify(params, cli.Params.BackKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		//签名校验失败
		reqJson, _ := json.Marshal(req)
		cli.logger.Errorf("ELKCrypto back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}

// 充值/提现都会调用这个
func (cli *Client) CurrencyCallback(req ELKCurrencyBackReq, processor func(ELKCurrencyBackReq) error) error {
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
		cli.logger.Errorf("ELKCur back verify fail, req: %s", string(reqJson))
		return errors.New("sign verify error")
	}

	//开始处理
	return processor(req)
}
