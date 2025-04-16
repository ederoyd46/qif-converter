package scanner

import (
	"bufio"
	"io"
	"qif-converter/model"
)

const separator = "^"

type entityType int

const (
	none entityType = iota
	account
)

// ScanAccounts reads a QIF file and returns a list of accounts.
func ScanAccounts(file io.Reader) model.Accounts {
	scanner := bufio.NewScanner(file)
	accountMap := make(map[string]*model.Account)
	linkedTransactions := make([]model.TransactionEntry, 0)

	var currentAccount *model.Account
	var currentEntityType entityType

	for scanner.Scan() {
		header := scanner.Text()
		switch header {
		case "!Account":
			currentAccount = model.NewAccount(model.ReadAccountEntry(scanner, separator))
			accountMap[currentAccount.GetAccountEntry().GetName()] = currentAccount
			currentEntityType = account
		case "!Type:Bank", "!Type:CCard":
			currentEntityType = account
		case "!Type:Class", "!Type:Cat":
			currentEntityType = none
		}

		switch currentEntityType {
		case account:
			transaction := model.ReadTransactionEntry(scanner, separator)
			currentAccount.AppendTransaction(transaction)
			if transaction.IsLinked() {
				linkedTransactions = append(linkedTransactions, transaction)
			}
		case none:
			skipEntry(scanner)
		}
	}

	populateLinkedTransaction(accountMap, linkedTransactions)

	accounts := make(model.Accounts, 0)

	for _, value := range accountMap {
		accounts = append(accounts, value)
	}

	return accounts
}

// populateLinkedTransaction add the linked transactions to their counterpart account so their amounts will total correctly
func populateLinkedTransaction(accountMap map[string]*model.Account, linkedTransactions []model.TransactionEntry) {
	for _, transaction := range linkedTransactions {
		if account, ok := accountMap[transaction.GetLinkedAccount()]; ok {
			transaction.SetAmount(transaction.GetAmount() * -1)
			account.AppendTransaction(transaction)
		}
	}
}

// skipEntry skips the current entry in the scanner.
func skipEntry(scanner *bufio.Scanner) {
	for {
		value := scanner.Text()
		if value == separator {
			break
		}
		scanner.Scan()
	}
}
