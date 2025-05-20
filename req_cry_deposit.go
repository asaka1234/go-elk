package go_elk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-elk/utils"
	"io/ioutil"
	"log"
	"net/http"
)

// 下单
func (cli *Client) CryDeposit(req ELKCryDepositReq) (*ELKCryDepositRsp, error) {

	rawURL := cli.CryDepositUrl

	// Convert struct to map for signing
	params := make(map[string]interface{})
	jsonData, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}
	if err := json.Unmarshal(jsonData, &params); err != nil {
		return nil, fmt.Errorf("failed to unmarshal to map: %v", err)
	}

	// Generate signature
	signature, err := utils.Sign(params, cli.AccessKey)
	if err != nil {
		return nil, fmt.Errorf("signature generation failed: %v", err)
	}
	req.Signature = signature

	// Prepare request JSON
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal signed request: %v", err)
	}
	log.Printf("ELKCryService#deposit#json: %s", jsonReq)

	// Send HTTP request
	resp, err := http.Post(rawURL, "application/json", bytes.NewBuffer(jsonReq))
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
	var result ELKCryDepositRsp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &result, nil
}
