package model_test

import (
	"bufio"
	"qif-converter/model"
	"reflect"
	"strings"
	"testing"
)

func TestReadAccountEntry_ValidEntry(t *testing.T) {
	t.Parallel()
	input := `!Account
NMy Test Bank Account
TBank
^`
	scanner := bufio.NewScanner(strings.NewReader(input))
	scanner.Split(bufio.ScanLines)
	scanner.Scan() // Need to call Scan once to initialize the scanner

	want := model.NewAccountEntry("My Test Bank Account", "Bank")
	got := model.ReadAccountEntry(scanner, "^")

	if !reflect.DeepEqual(got, want) {
		t.Errorf("ReadAccountEntry() returned incorrect AccountEntry, got %+v, expected %+v", got, want)
	}
}

//TODO More tests needed
