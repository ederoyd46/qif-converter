package model

import (
	"bufio"
	"encoding/json"
)

type AccountEntry struct {
	name        string //N
	accountType string //T
}

func (self *AccountEntry) SetName(name string) {
	self.name = name
}

func (self AccountEntry) GetName() string {
	return self.name
}

func (self AccountEntry) GetAccountType() string {
	return self.accountType
}

func (self *AccountEntry) SetAccountType(accountType string) {
	self.accountType = accountType
}

func NewAccountEntry(name string, accountType string) AccountEntry {
	return AccountEntry{
		name:        name,
		accountType: accountType,
	}
}

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
