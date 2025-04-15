package model

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

// TransactionEntry transaction container
type TransactionEntry struct {
	date     string  //D
	amount   float32 //T | U
	memo     string  //M
	payee    string  //P
	category string  //L
}

// NewTransactionEntry Creates a new transaction entry
func NewTransactionEntry(date string, amount float32, memo string, payee string, category string) TransactionEntry {
	return TransactionEntry{
		date:     date,
		amount:   amount,
		memo:     memo,
		payee:    payee,
		category: category,
	}
}

// SetDate sets the date of the transaction entry
func (self *TransactionEntry) SetDate(date string) {
	self.date = date
}

// GetDate gets the date of the transaction entry
func (self TransactionEntry) GetDate() string {
	return self.date
}

// SetAmount sets the amount of the transaction entry
func (self *TransactionEntry) SetAmount(amount float32) {
	self.amount = amount
}

// GetAmount gets the amount of the transaction entry
func (self TransactionEntry) GetAmount() float32 {
	return self.amount
}

// SetMemo sets the memo of the transaction entry
func (self *TransactionEntry) SetMemo(memo string) {
	if memo != "(null)" {
		self.memo = memo
	}
}

// GetMemo gets the memo of the transaction entry
func (self TransactionEntry) GetMemo() string {
	return self.memo
}

// SetPayee sets the payee of the transaction entry
func (self *TransactionEntry) SetPayee(payee string) {
	self.payee = payee
}

// GetPayee gets the payee of the transaction entry
func (self TransactionEntry) GetPayee() string {
	return self.payee
}

// SetCategory sets the category of the transaction entry
func (self *TransactionEntry) SetCategory(category string) {
	self.category = category
}

// GetCategory gets the category of the transaction entry
func (self TransactionEntry) GetCategory() string {
	return self.category
}

// ReadTransactionEntry Read buffer to build a transaction entry
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
			entry.SetDate(val)
			break
		case "T":
			amount, err := strconv.ParseFloat(val, 32)
			if err != nil {
				os.Stderr.WriteString(fmt.Sprintf("Could not convert value %q to a number\n", val))
				amount = 0
			}
			entry.SetAmount(float32(amount))
			break
		case "M":
			entry.SetMemo(val)
			break
		case "P":
			entry.SetPayee(val)
			break
		case "L":
			entry.SetCategory(val)
		}

		more := scanner.Scan()
		if !more {
			break
		}
	}
	return entry
}

// MarshalJSON mashal TransactionEntry to publicly exportable format
func (self TransactionEntry) MarshalJSON() ([]byte, error) {
	type PublicTransactionEntry struct {
		Date     string  `json:"date"`
		Amount   float32 `json:"amount"`
		Memo     string  `json:"memo"`
		Payee    string  `json:"payee"`
		Category string  `json:"category"`
	}

	entry := PublicTransactionEntry{
		Date:     self.date,
		Amount:   self.amount,
		Memo:     self.memo,
		Payee:    self.payee,
		Category: self.category,
	}

	return json.Marshal(entry)
}

// String document public functions
func (self TransactionEntry) String() string {
	return fmt.Sprintf("TransactionEntry{date=%s, amount=%.2f, memo=%s, payee=%s, category=%s}", self.date, self.amount, self.memo, self.payee, self.category)
}
