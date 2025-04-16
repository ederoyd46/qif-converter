package model_test

import (
	"bufio"
	"qif-converter/model"
	"reflect"
	"strings"
	"testing"
)

func TestReadTransactionEntry_ValidEntry(t *testing.T) {
	t.Parallel()
	input := `D2023-10-26
T100.50
MTest Memo
PTest Payee
LTest Category
^`
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	scanner.Scan() // Need to call Scan once to initialize the scanner

	var want model.TransactionEntry
	want.SetDate("2023-10-26")
	want.SetAmount(100.50)
	want.SetMemo("Test Memo")
	want.SetPayee("Test Payee")
	want.SetCategory("Test Category")

	got := model.ReadTransactionEntry(scanner, "^")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadTransactionEntry() returned incorrect TransactionEntry, got %+v, expected %+v", got, want)
	}
}

//TODO More tests needed
