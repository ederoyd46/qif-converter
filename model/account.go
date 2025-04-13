package model

type Accounts = []*Account

type Account struct {
	AccountEntry `json:"account"`
	Transactions []TransactionEntry `json:"transactions"`
}
