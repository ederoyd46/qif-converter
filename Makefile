.PHONY: build

build:
	go build .

release:
	go build -ldflags "-w" -o target/qif-converter .

release-all-architectures:
	GOOS='linux' GOARCH='amd64' go build -ldflags "-w" -o target/qif-converter-linux-amd64 .
	GOOS='linux' GOARCH='arm64' go build -ldflags "-w" -o target/qif-converter-linux-arm .
	GOOS='darwin' GOARCH='amd64' go build -ldflags "-w" -o target/qif-converter-darwin-amd64 .
	GOOS='darwin' GOARCH='arm64' go build -ldflags "-w" -o target/qif-converter-darwin-arm .
	GOOS='windows' GOARCH='amd64' go build -ldflags "-w" -o target/qif-converter-windows-amd64.exe .
	GOOS='windows' GOARCH='arm64' go build -ldflags "-w" -o target/qif-converter-windows-arm.exe .

test_homebank:
	go run main.go homebank.qif | duckdb transactions.duckdb -c "create table homebank_transactions as select unnest(account), unnest(transactions, recursive := true) from read_json('/dev/stdin')"

test_homebank_results:
	duckdb transactions.duckdb -c "select account_name, count(*) as num_of_transactions, round(sum(amount),2) as balance from homebank_transactions group by account_name order by balance desc"

test_moneydance:
	go run main.go moneydance.qif | duckdb transactions.duckdb -c "create table moneydance_transactions as select unnest(account), unnest(transactions, recursive := true) from read_json('/dev/stdin')"

test_moneydance_results:
	duckdb transactions.duckdb -c "select account_name, count(*) as num_of_transactions, round(sum(amount),2) as balance from moneydance_transactions group by account_name order by balance desc"
