package go_elk

import (
	"encoding/json"
	"github.com/asaka1234/go-elk/utils"
	"log"
)

// 充值/提现的回调处理(传入一个处理函数)
func (cli *Client) CryWithdrawCallback(req ELKCryWithdrawBackReq, processor func(ELKCryWithdrawBackReq) error) error {
	//验证签名
	//验证签名
	params := make(map[string]interface{})
	jsonData, err := json.Marshal(req)
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
		return err
	}
	if err := json.Unmarshal(jsonData, &params); err != nil {
		log.Printf("JSON unmarshal error: %v", err)
		return err
	}

	// Verify signature
	flag, err := utils.Verify(params, cli.BackKey)
	if err != nil {
		log.Printf("Signature verification error: %v", err)
		return err
	}
	if !flag {
		reqJson, _ := json.Marshal(req)
		log.Printf("ELKCur back verify fail, req: %s", string(reqJson))
	}

	//开始处理
	return processor(req)
}
