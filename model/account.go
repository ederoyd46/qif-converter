package model

import "encoding/json"

type Accounts = []*Account

type Account struct {
	accountEntry AccountEntry
	transactions []TransactionEntry
}

// NewAccount create new account and return as pointer
func NewAccount(accountEntry AccountEntry) *Account {
	return &Account{
		accountEntry: accountEntry,
		transactions: []TransactionEntry{},
	}
}

// SetAccountEntry sets the account entry for the account.
func (self *Account) SetAccountEntry(accountEntry AccountEntry) {
	self.accountEntry = accountEntry
}

// GetAccountEntry returns the account entry for the account.
func (self Account) GetAccountEntry() AccountEntry {
	return self.accountEntry
}

// SetTransactions sets the transactions for the account.
func (self *Account) SetTransactions(transactions []TransactionEntry) {
	self.transactions = transactions
}

// GetTransactions returns the transactions for the account.
func (self Account) GetTransactions() []TransactionEntry {
	return self.transactions
}

// AppendTransaction appends a transaction to the account.
func (self *Account) AppendTransaction(transaction TransactionEntry) {
	self.transactions = append(self.transactions, transaction)
}

// MarshalJSON marshals the account to JSON.
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
