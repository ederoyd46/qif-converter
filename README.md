# QIF Converter

This project is a Go-based application for converting QIF ([Quicken Interchange Format](https://en.wikipedia.org/wiki/Quicken_Interchange_Format)) files.

Currently only exported qif files from Homebank are supported. But it should work to some degree with any qif file.

## Prerequisites

- Go 1.20 or later installed on your system.
- Git installed for version control.

## Getting Started

1. Clone the repository:
    ```bash
    git clone https://github.com/ederoyd46/qif-converter.git
    cd qif-converter
    ```

2. Build the project:
    ```bash
    go build
    ```

3. Run the application:
    ```bash
    ./qif-converter accounts.qif
    ```

## Testing

Run the tests using:
```bash
go test ./...
```

## Contributing

1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Commit your changes and push the branch.
4. Open a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

## Issues

* Currently this does not handle (linked) transfers when exporting all accounts from Homebank as the data is not in the expected location. *As a work around, export the accounts as individual qif files.*

## Nix Alternative Getting Started
As an alternative to the above, you can also build using the [Nix](https://nixos.org/) package manager.

1. Clone the repository:
    ```bash
    git clone https://github.com/ederoyd46/qif-converter.git
    cd qif-converter
    ```

2 Build using Nix
  ```bash
  nix build .#build
  ```

3 Run using Nix
  ```bash
  nix run .#default -- accounts.qif
  ```

4 Build Docker Image (for linux arm)
  ```bash
  nix build .#packages.aarch64-linux.buildDockerImage
  ```

5. Load Docker Image
```bash
docker load < result
```

# DuckDB - Import

#### Import data into a table

```bash
./qif-converter homebank.qif | duckdb transactions.duckdb -c "create table homebank_transactions as select unnest(account), unnest(transactions, recursive := true) from read_json('/dev/stdin')"
```

#### Number of transactions per account and balance
```sql
select  account_name,
        count(*) as num_of_transactions,
        round(sum(amount),2) as balance
from homebank_transactions
group by account_name
order by balance desc;
```
