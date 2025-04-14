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

	want := model.NewTransactionEntry("2023-10-26", 100.50, "Test Memo", "Test Payee", "Test Category")
	got := model.ReadTransactionEntry(scanner, "^")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadTransactionEntry() returned incorrect TransactionEntry, got %+v, expected %+v", got, want)
	}
}

//TODO More tests needed
