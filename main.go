package main

import (
	"fmt"
	"os"
	"qif-converter/scanner"
	"qif-converter/transformer"
	"strings"
)

func main() {
	fileName, err := getFileName(os.Args)
	handleFatalError(err)

	file, err := os.Open(fileName)
	handleFatalError(err)
	defer file.Close()

	accounts := scanner.ScanAccounts(file)

	result, err := transformer.ToJSON(accounts)
	handleFatalError(err)
	os.Stdout.Write(result)
}

// Get the file name from the list of arguments, it does not matter which position it's in.
func getFileName(args []string) (string, error) {
	for _, arg := range args {
		if strings.HasSuffix(arg, ".qif") {
			return arg, nil
		}
	}

	return "", fmt.Errorf("Could not find file in args %v", args)
}

// Default error handler for fatal errors.
func handleFatalError(err error) {
	if err != nil {
		os.Stderr.WriteString(fmt.Sprintf("Fatal error occurred [%v]", err))
		os.Exit(1)
	}
}
