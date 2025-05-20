package go_elk

// 充值/提现的回调处理(传入一个处理函数)
func (cli *Client) CryWithdrawCallback(req ELKCryWithdrawBackReq, processor func(ELKCryWithdrawBackReq) error) error {
	//验证签名
	//TODO

	//开始处理
	return processor(req)
}
