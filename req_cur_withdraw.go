package go_elk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-elk/utils"
	"io/ioutil"
	"net/http"
)

// 下单(充值/提现是同一个接口)
func (cli *Client) CurWithdraw(req ELKCurWithdrawReq) (*ELKCurWithdrawRsp, error) {

	rawURL := cli.CurWithdrawUrl

	// Convert struct to map for signing
	params := make(map[string]interface{})
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %v", err)
	}
	if err := json.Unmarshal(jsonData, &params); err != nil {
		return nil, fmt.Errorf("JSON unmarshal error: %v", err)
	}

	// Generate signature
	signature, err := utils.Sign(params, cli.AccessKey)
	if err != nil {
		return nil, fmt.Errorf("signature generation error: %v", err)
	}
	req.Signature = signature

	// Prepare request
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("JSON marshal error: %v", err)
	}
	cli.logger.Infof("ELKCurService#withdraw#json: %s", jsonReq)

	// Send HTTP request
	resp, err := http.Post(rawURL, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("HTTP request error: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read response error: %v", err)
	}

	// Parse response
	var result ELKCurWithdrawRsp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("parse response error: %v", err)
	}

	return &result, nil
}
