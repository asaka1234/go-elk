package go_elk

func (cli *Client) CryDepositCallback(req ELKCryDepositBackReq, processor func(ELKCryDepositBackReq) error) error {
	//验证签名
	//TODO

	//开始处理
	return processor(req)
}

func (cli *Client) CurDepositCallback(req ELKCurDepositBackReq, processor func(ELKCurDepositBackReq) error) error {
	//验证签名
	//TODO

	//开始处理
	return processor(req)
}
