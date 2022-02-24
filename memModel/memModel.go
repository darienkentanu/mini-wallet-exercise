package memModel

var Wallets = []*MyWallet{}
var ReferencesIDS = make(map[string]bool)

func init() {
	Wallets = append(Wallets, &MyWallet{"50535246-dcb2-4929-8cc9-004ea06f5241", "ea0212d3-abd6-406f-8c67-868e814a2436", "disabled", "", 0})
}

type POSTCustomer struct {
	CustomerXID string `form:"customer_xid"`
}

type MyWallet struct {
	Id       string `json:"id"`
	OwnedBy  string `json:"owned_by"`
	Status   string `json:"status"`
	EnableAt string `json:"enable_at"`
	Balance  int    `json:"balance"`
}

type POSTAddVirtualMoney struct {
	Amount      int    `form:"amount"`
	ReferenceID string `form:"reference_id"`
}

type POSTWithdrawVirtualMoney struct {
	Amount      int    `form:"amount"`
	ReferenceID string `form:"reference_id"`
}

type AddMoney struct {
	ID          string `json:"id"`
	DepositedBy string `json:"deposited_by"`
	Status      string `json:"status"`
	DepositedAt string `json:"deposited_at"`
	Amount      int    `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

type Withdrawal struct {
	ID          string `json:"id"`
	WithdrawnBy string `json:"withdrawn_by"`
	Status      string `json:"status"`
	WithdrawnAt string `json:"withdrawn_at"`
	Amount      int    `json:"amount"`
	ReferenceID string `json:"reference_id"`
}

type PATCHdisable struct {
	IsDisabled string `form:"is_disabled"`
}
