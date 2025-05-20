package utils

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"
)

func Sign(params map[string]interface{}, key string) (string, error) {
	// 1. Validate key
	if key == "" {
		return "", errors.New("APP_KEY 参数为空，请填写")
	}

	// 2. Get and sort keys
	keys := make([]string, 0, len(params))
	for k := range params {
		keys = append(keys, k)
	}
	sort.Strings(keys) // ASCII ascending order

	// 3. Build sign string
	var sb strings.Builder
	for _, k := range keys {
		sb.WriteString(fmt.Sprintf("%s=%v&", k, params[k]))
	}
	sb.WriteString(fmt.Sprintf("key=%s", key))

	// 4. Generate MD5
	signStr := sb.String()
	hash := md5.Sum([]byte(signStr))
	signResult := hex.EncodeToString(hash[:])

	// Debug print (optional)
	fmt.Printf("验签str: %s结果: %s\n", signStr, signResult)

	return signResult, nil
}

func Verify(params map[string]interface{}, signKey string) (bool, error) {
	// Check if signature exists in params
	signature, exists := params["signature"]
	if !exists {
		return false, nil
	}

	// Remove signature from params for verification
	delete(params, "signature")

	// Generate current signature
	currentSignature, err := Sign(params, signKey)
	if err != nil {
		return false, fmt.Errorf("signature generation failed: %w", err)
	}

	// Compare signatures
	return signature.(string) == currentSignature, nil
}
