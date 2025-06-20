package go_elk

import (
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestCryDeposit(t *testing.T) {
	vLog := VLog{}
	//构造client
	cli := NewClient(vLog, &ELKInitParams{MerchantInfo{MERCHANT_ID, ACCESS_KEY, BACK_KEY}, CUR_DEPOSIT_URL, CUR_WITHDRAW_URL, CRY_DEPOSIT_URL, CRY_WITHDRAW_URL})

	//发请求
	resp, err := cli.CryDeposit(GenCryDepositRequestDemo())
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	cli.logger.Infof("resp:%+v\n", resp)
}

func GenCryDepositRequestDemo() ELKCryDepositReq {
	return ELKCryDepositReq{
		UniqueCode: uuid.New().String(),
		Protocol:   "TRC20",
		CoinName:   "USDT",
		Amount:     "600.00",
		OrderId:    "1234",
	}
}
