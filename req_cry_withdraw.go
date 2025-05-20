package go_elk

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/asaka1234/go-elk/utils"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func (cli *Client) CryWithdraw(req ELKCryWithdrawReq) (*ELKCryWithdrawRsp, error) {

	rawURL := cli.CurWithdrawUrl

	// 2. Convert struct to map for signing
	params := make(map[string]interface{})
	if err := mapstructure.Decode(req, &params); err != nil {
		return nil, fmt.Errorf("failed to convert request to map: %v", err)
	}

	// 3. Generate signature
	signature, err := utils.Sign(params, cli.AccessKey)
	if err != nil {
		return nil, fmt.Errorf("signature generation failed: %v", err)
	}
	req.Signature = signature

	// 4. Prepare request
	jsonReq, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal request: %v", err)
	}
	log.Printf("ELKCryService#withdraw#json: %s", jsonReq)

	// 5. Send HTTP request
	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Post(rawURL, "application/json", bytes.NewBuffer(jsonReq))
	if err != nil {
		return nil, fmt.Errorf("HTTP request failed: %v", err)
	}
	defer resp.Body.Close()

	// 6. Handle response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	// 7. Parse response
	var result ELKCryWithdrawRsp
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, fmt.Errorf("failed to parse response: %v", err)
	}

	return &result, nil
}
