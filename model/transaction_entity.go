package model

type TransactionEntry struct {
	Date     string  `json:"date"`     //D
	Amount   float32 `json:"amount"`   //T | U
	Memo     string  `json:"memo"`     //M
	Payee    string  `json:"payee"`    //P
	Cleared  bool    `json:"cleared"`  //C
	Category string  `json:"category"` //L
}
