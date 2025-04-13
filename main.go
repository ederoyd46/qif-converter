package main

import (
	"bufio"
	"fmt"
	"os"
	"qif-converter/model"
	"qif-converter/model/transformer"
	"strconv"
	"strings"
)

// type Output int
// const (
// 	OutputJSON Output = iota
// 	OuputCSV
// )

const EndOfRecord = "^"

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

	for scanner.Scan() {
		header := scanner.Text()
		switch header {
		case "!Account":
			//New Account
			account := model.Account{AccountEntry: readAccountEntry(scanner), Transactions: []model.TransactionEntry{}}
			accounts = append(accounts, &account)
			currentAccount = &account
			break
		default:
			//Probably a transaction
			currentAccount.Transactions = append(currentAccount.Transactions, readTransactionEntry(scanner))
		}
	}

	result, err := transformer.ToJSON(accounts)
	handleError(err)
	os.Stdout.Write(result)

}

func readAccountEntry(scanner *bufio.Scanner) model.AccountEntry {
	entry := model.AccountEntry{}
	//Move to the next line as we don't care about the !Account header
	scanner.Scan()
	for {
		value := scanner.Text()
		if value == EndOfRecord {
			break
		}

		key := value[0:1]
		val := value[1:]

		switch key {
		case "N":
			entry.Name = val
			break
		case "T":
			entry.AccountType = val
		}
		more := scanner.Scan()
		if !more {
			break
		}
	}
	return entry
}

func readTransactionEntry(scanner *bufio.Scanner) model.TransactionEntry {
	entry := model.TransactionEntry{}
	for {
		value := scanner.Text()
		if value == EndOfRecord {
			break
		}

		key := value[0:1]
		val := value[1:]

		switch key {
		case "D":
			entry.Date = val
			break
		case "T":
			amount, err := strconv.ParseFloat(val, 32)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("Could not convert value %q to a number\n", val))
				amount = 0
			}
			entry.Amount = float32(amount)
			break
		case "M":
			entry.Memo = val
			break
		case "P":
			entry.Payee = val
			break
		case "L":
			entry.Category = val
		}

		more := scanner.Scan()
		if !more {
			break
		}
	}
	return entry
}
