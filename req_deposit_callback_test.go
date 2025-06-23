package go_elk

import (
	"encoding/json"
	"fmt"
	"testing"
)

type VLog struct {
}

func (l VLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
func (l VLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

func TestCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ELKInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, CUR_DEPOSIT_URL, CUR_WITHDRAW_URL, CRY_DEPOSIT_URL, CRY_WITHDRAW_URL})

	//1. 获取请求
	req := GenCallbackRequestDemo() //提现的返回
	var backReq ELKCurrencyBackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	//2. 处理请求
	err = cli.CurrencyCallback(backReq, func(ELKCurrencyBackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", backReq)
}

func GenCallbackRequestDemo() string {
	return `{
	"apiOrderNo": "202506191258180899",
	"amount": "10.8815",
	"money": "79",
	"uniqueCode": "caaf91c5-3765-4527-a777-400217267e5e",
	"signature": "5f994a7e08ee73f1c07efabd7e9c0641",
	"tradeStatus": "1",
	"tradeId": "814309903453126656"
}`
}

func TestCryptoCallback(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ELKInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, CUR_DEPOSIT_URL, CUR_WITHDRAW_URL, CRY_DEPOSIT_URL, CRY_WITHDRAW_URL})

	//1. 获取请求
	req := GenCallbackCryptoRequestDemo() //提现的返回
	var backReq ELKCryDepositBackReq
	err := json.Unmarshal([]byte(req), &backReq)
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}

	//2. 处理请求
	err = cli.CryDepositCallback(backReq, "aaa", func(ELKCryDepositBackReq) error { return nil })
	if err != nil {
		cli.logger.Errorf("Error:%s", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", backReq)
}

// usdt充值回调
func GenCallbackCryptoRequestDemo() string {
	return `{
	"apiOrderNo": "202506230615240440",
	"amount": "0.9500",
	"protocol": "TRC20",
	"orderAmount": "1.0000",
	"uniqueCode": "5553768",
	"fee": "0.0500",
	"txId": "13ea2909937520be8777c1c4ba0a2aa450b6ebcf5c31c3940d52d1e2a0753913",
	"coinName": "USDT",
	"tradeId": 815658084854136832
}`
}
