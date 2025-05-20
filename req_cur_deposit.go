package go_elk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-elk/utils"
	"io/ioutil"
	"net/http"
)

// 下单
func (cli *Client) CurDeposit(req ELKCurDepositReq) (*ELKCurDepositRsp, error) {

	rawURL := cli.CurDepositUrl

	// Log request
	cli.logger.Infof("ELKCurService#req: %+v", req)

	// Prepare params
	params := map[string]interface{}{
		"uid":        req.UID,
		"uniqueCode": req.UniqueCode,
		"money":      req.Money,
		"payType":    req.PayType,
		"orderId":    req.OrderId,
		"payerName":  req.PayerName,
	}

	// Add jumpUrl if not empty
	if req.JumpUrl != "" {
		params["jumpUrl"] = req.JumpUrl
	}

	// Generate signature
	signature, err := utils.Sign(params, cli.AccessKey)
	if err != nil {
		return nil, fmt.Errorf("signature generation failed: %v", err)
	}
	params["signature"] = signature

	// Convert to JSON
	jsonStr, err := json.Marshal(params)
	if err != nil {
		return nil, fmt.Errorf("JSON marshaling failed: %v", err)
	}
	cli.logger.Infof("ELKCurService#deposit#json: %s", jsonStr)

	// Send HTTP request
	resp, err := http.Post(rawURL, "application/json", bytes.NewBuffer(jsonStr))
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response: %v", err)
	}

	// Parse response
	var result ELKCurDepositRsp
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &result, nil
}
