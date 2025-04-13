package model

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type TransactionEntry struct {
	Date     string  `json:"date"`     //D
	Amount   float32 `json:"amount"`   //T | U
	Memo     string  `json:"memo"`     //M
	Payee    string  `json:"payee"`    //P
	Category string  `json:"category"` //L
}

func ReadTransactionEntry(scanner *bufio.Scanner, recordSeparator string) TransactionEntry {
	entry := TransactionEntry{}
	for {
		value := scanner.Text()
		if value == recordSeparator {
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
