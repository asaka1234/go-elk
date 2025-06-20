package go_elk

import (
	"testing"
)

func TestCryWithdraw(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ELKInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, CUR_DEPOSIT_URL, CUR_WITHDRAW_URL, CRY_DEPOSIT_URL, CRY_WITHDRAW_URL})

	//发请求
	resp, err := cli.CryWithdraw(GenCryWithdrawRequestDemo())
	if err != nil {
		cli.logger.Errorf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenCryWithdrawRequestDemo() ELKCryWithdrawReq {
	return ELKCryWithdrawReq{
		ToAddress: "hiahaihia",
		ChainName: "TRC20",
		CoinName:  "USDT",
		Amount:    "600.00",
		OrderId:   "123",
	}
}
