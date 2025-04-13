package main

import (
	"bufio"
	"fmt"
	"os"
	"qif-converter/model"
	"qif-converter/model/transformer"
	"strings"
)

const EndOfRecord = "^"

type EntityType int

const (
	None EntityType = iota
	Account
)

func getFileName(args []string) (string, error) {
	for _, arg := range args {
		if strings.HasSuffix(arg, ".qif") {
			return arg, nil
		}
	}

	return "", fmt.Errorf("Could not find file in args %v", args)
}

func handleError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	fileName, err := getFileName(os.Args)
	handleError(err)

	file, err := os.Open(fileName)
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	accounts := make(model.Accounts, 0)
	var currentAccount *model.Account
	var currentEntity EntityType

	for scanner.Scan() {
		header := scanner.Text()
		switch header {
		case "!Account":
			account := model.Account{AccountEntry: model.ReadAccountEntry(scanner, EndOfRecord), Transactions: []model.TransactionEntry{}}
			accounts = append(accounts, &account)
			currentAccount = &account
			currentEntity = Account
			break
		case "!Type:Class":
			currentEntity = None
			break
		case "!Type:Bank":
			currentEntity = Account
			break
		case "!Type:CCard":
			currentEntity = Account
			break
		case "!Type:Cat":
			currentEntity = None
			break
		}

		switch currentEntity {
		case Account:
			currentAccount.Transactions = append(currentAccount.Transactions, model.ReadTransactionEntry(scanner, EndOfRecord))
			break
		default:
			skipEntry(scanner)
		}
	}

	result, err := transformer.ToJSON(accounts)
	handleError(err)
	os.Stdout.Write(result)
}

func skipEntry(scanner *bufio.Scanner) {
	for {
		value := scanner.Text()
		if value == EndOfRecord {
			break
		}
		scanner.Scan()
	}
}
