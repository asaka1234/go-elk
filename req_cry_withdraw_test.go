package go_elk

import (
	"fmt"
	"testing"
)

func TestCryWithdraw(t *testing.T) {

	//构造client
	cli := NewClient(nil, ELKInitParams{MERCHANT_ID, ACCESS_KEY, BACK_KEY, CUR_DEPOSIT_URL, CUR_WITHDRAW_URL, CRY_DEPOSIT_URL, CRY_WITHDRAW_URL})

	//发请求
	resp, err := cli.CryWithdraw(GenCryWithdrawRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("resp:%+v\n", resp)
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
