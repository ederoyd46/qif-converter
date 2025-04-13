package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// type Output int
// const (
// 	OutputJSON Output = iota
// 	OuputCSV
// )

const EndOfRecord = "^"

type TransactionEntry struct {
	Date     string  `json:"date"`     //D
	Amount   float32 `json:"amount"`   //T | U
	Memo     string  `json:"memo"`     //M
	Payee    string  `json:"payee"`    //P
	Cleared  bool    `json:"cleared"`  //C
	Category string  `json:"category"` //L
}

type AccountEntry struct {
	Name        string `json:"name"` //N
	AccountType string `json:"type"` //T
}

type Account struct {
	AccountEntry
	Transactions []TransactionEntry
}

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

	accounts := make([]*Account, 0)
	var currentAccount *Account

	for scanner.Scan() {
		header := scanner.Text()
		switch header {
		case "!Account":
			//New Account
			account := Account{readAccountEntry(scanner), []TransactionEntry{}}
			accounts = append(accounts, &account)
			currentAccount = &account
			break
		default:
			currentAccount.Transactions = append(currentAccount.Transactions, readTransactionEntry(scanner))
		}
	}

	result, err := json.Marshal(accounts)
	handleError(err)
	os.Stdout.Write(result)

}

func readAccountEntry(scanner *bufio.Scanner) AccountEntry {
	entry := AccountEntry{}
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

func readTransactionEntry(scanner *bufio.Scanner) TransactionEntry {
	entry := TransactionEntry{}
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
				os.Stderr.WriteString(fmt.Sprintf("BUG! Could not convert value %q to a number\n", val))
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
