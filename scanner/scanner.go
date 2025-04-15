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
	accounts := make(model.Accounts, 0)
	var currentAccount *model.Account
	var currentEntity entityType

	for scanner.Scan() {
		header := scanner.Text()
		switch header {
		case "!Account":
			currentAccount = model.NewAccount(model.ReadAccountEntry(scanner, separator))
			accounts = append(accounts, currentAccount)
			currentEntity = account
		case "!Type:Bank", "!Type:CCard":
			currentEntity = account
		case "!Type:Class", "!Type:Cat":
			currentEntity = none
		}

		switch currentEntity {
		case account:
			currentAccount.AppendTransaction(model.ReadTransactionEntry(scanner, separator))
		case none:
			skipEntry(scanner)
		}
	}

	return accounts
}

// SkipEntry skips the current entry in the scanner.
func skipEntry(scanner *bufio.Scanner) {
	for {
		value := scanner.Text()
		if value == separator {
			break
		}
		scanner.Scan()
	}
}
