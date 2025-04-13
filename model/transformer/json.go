package transformer

import (
	"encoding/json"
	"qif-converter/model"
)

func ToJSON[T model.Accounts](t T) ([]byte, error) {
	return json.Marshal(t)
}
