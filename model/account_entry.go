package model

import (
	"bufio"
	"encoding/json"
	"fmt"
)

// AccountEntry account container
type AccountEntry struct {
	name        string //N
	accountType string //T
}

// SetName sets the name of the account entry
func (self *AccountEntry) SetName(name string) {
	self.name = name
}

// GetName returns the name of the account entry
func (self AccountEntry) GetName() string {
	return self.name
}

// SetAccountType sets the account type of the account entry
func (self *AccountEntry) SetAccountType(accountType string) {
	self.accountType = accountType
}

// GetAccountType returns the name of the account entry
func (self AccountEntry) GetAccountType() string {
	return self.name
}

// NewAccountEntry creates a new account entry
func NewAccountEntry(name string, accountType string) AccountEntry {
	return AccountEntry{
		name:        name,
		accountType: accountType,
	}
}

// ReadAccountEntry Read buffer to build an account entry
func ReadAccountEntry(scanner *bufio.Scanner, recordSeparator string) AccountEntry {
	entry := AccountEntry{}
	//Move to the next line as we don't care about the !Account header
	for scanner.Scan() {
		value := scanner.Text()
		if value == recordSeparator {
			break
		}

		key := value[0:1]
		val := value[1:]

		switch key {
		case "N":
			entry.SetName(val)
		case "T":
			entry.SetAccountType(val)
		}
	}
	return entry
}

// MarshalJSON Mashal AccountEntry to publicly exportable format
func (self AccountEntry) MarshalJSON() ([]byte, error) {
	type PublicAccountEntry struct {
		Name        string `json:"account_name"`
		AccountType string `json:"account_type"`
	}

	entry := PublicAccountEntry{
		Name:        self.name,
		AccountType: self.accountType,
	}

	return json.Marshal(entry)
}

// String document public functions
func (self AccountEntry) String() string {
	return fmt.Sprintf("AccountEntry{name=%s, accountTyoe=%s}", self.name, self.accountType)
}
