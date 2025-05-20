package go_elk

type ELKCurDepositReq struct {
	UID        int    `json:"uid"`
	UniqueCode string `json:"uniqueCode"`
	Money      string `json:"money"`
	PayType    int    `json:"payType"` // 1: UniPay 2: Alipay 3: WeChat
	OrderId    string `json:"orderId"`
	Signature  string `json:"signature"`
	PayerName  string `json:"payerName"`
	JumpUrl    string `json:"jumpUrl"`
}

type ELKCurDepositRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

type ELKCurWithdrawReq struct {
	UID       int    `json:"uid"`
	Money     string `json:"money"`
	OrderId   string `json:"orderId"`
	Signature string `json:"signature"`
	PayerName string `json:"payerName"`
	CardNo    string `json:"cardNo"`
	BankName  string `json:"bankName"`
}

type ELKCurWithdrawRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

//--------------------------------------------------

type ELKCryDepositReq struct {
	UID        int    `json:"uid"`
	UniqueCode string `json:"uniqueCode"`
	Protocol   string `json:"protocol"`
	CoinName   string `json:"coinName"`
	OrderId    string `json:"orderId"`
	Amount     string `json:"amount"`
	Signature  string `json:"signature"`
}

type ELKCryDepositRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

type ELKCryWithdrawReq struct {
	UID       int    `json:"uid"`
	ChainName string `json:"chainName"`
	CoinName  string `json:"coinName"`
	OrderId   string `json:"orderId"`
	Amount    string `json:"amount"`
	ToAddress string `json:"toAddress"`
	Signature string `json:"signature"`
}

type ELKCryWithdrawRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

// --------------------------------------------------

type ELKCurDepositBackReq struct {
	ApiOrderNo        string `json:"apiOrderNo"`
	Money             string `json:"money"`
	TradeStatus       int    `json:"tradeStatus"`
	TradeId           string `json:"tradeId"`
	WithdrawalOrderNo string `json:"withdrawalOrderNo"`
	UniqueCode        string `json:"uniqueCode"`
	Signature         string `json:"signature"`
}

type ELKCurDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

//--------------------------------------------------

type ELKCryDepositBackReq struct {
	ApiOrderNo  string `json:"apiOrderNo"`
	TradeId     string `json:"tradeId"`
	TxId        string `json:"txId"`
	UniqueCode  string `json:"uniqueCode"`
	Protocol    string `json:"protocol"`
	CoinName    string `json:"coinName"`
	Amount      string `json:"amount"`
	OrderAmount string `json:"orderAmount"`
	Fee         string `json:"fee"`
}

type ELKCryDepositBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}

//--------------------------------------------------

type ELKCryWithdrawBackReq struct {
	Amount      string `json:"amount"`
	Fee         string `json:"fee"`
	CoinName    string `json:"coinName"`
	TradeId     string `json:"tradeId"`
	ApiOrderNo  string `json:"apiOrderNo"`
	TradeStatus int    `json:"tradeStatus"`
	TxId        string `json:"txId"`
	ToAddress   string `json:"toAddress"`
	Signature   string `json:"signature"`
}

type ELKCryWithdrawBackRsp struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    string `json:"data"`
	Success bool   `json:"success"`
}
