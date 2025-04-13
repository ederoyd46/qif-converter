package model

import "bufio"

type AccountEntry struct {
	Name        string `json:"name"` //N
	AccountType string `json:"type"` //T
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
