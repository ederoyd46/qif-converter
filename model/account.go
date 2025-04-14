package model

import "encoding/json"

type Accounts = []*Account

type Account struct {
	accountEntry AccountEntry
	transactions []TransactionEntry
}

func NewAccount(accountEntry AccountEntry) Account {
	return Account{
		accountEntry: accountEntry,
		transactions: []TransactionEntry{},
	}
}

func (self *Account) SetAccountEntry(accountEntry AccountEntry) {
	self.accountEntry = accountEntry
}

func (self Account) GetAccountEntry() AccountEntry {
	return self.accountEntry
}

func (self *Account) SetTransactions(transactions []TransactionEntry) {
	self.transactions = transactions
}

func (self Account) GetTransactions() []TransactionEntry {
	return self.transactions
}

func (self *Account) AppendTransaction(transaction TransactionEntry) {
	self.transactions = append(self.transactions, transaction)
}

func (self Account) MarshalJSON() ([]byte, error) {
	type PublicAccount struct {
		AccountEntry AccountEntry       `json:"account"`
		Transactions []TransactionEntry `json:"transactions"`
	}

	entry := PublicAccount{
		AccountEntry: self.accountEntry,
		Transactions: self.transactions,
	}

	return json.Marshal(entry)
}
