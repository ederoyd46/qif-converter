package transformer

import (
	"encoding/json"
	"qif-converter/model"
)

// ToJSON converts a list of accounts to JSON format. We _only_ want to allow serialising Accounts
func ToJSON[T model.Accounts](t T) ([]byte, error) {
	return json.Marshal(t)
}
