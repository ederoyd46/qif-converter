# QIF Converter

This project is a Go-based application for converting QIF (Quicken Interchange Format) files.

Currently on exported qif files from Homebank are supported.

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
    ./qif-converter {accounts.qif}
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

# DuckDB - Import

### Import data into a table
```sql
create table homebank_transactions as select unnest(account), unnest(transactions, recursive := true) from read_json(homebank.json);
```

### Number of transactions per account and balance
```sql
select  name, 
        count(*) as num_of_transactions, 
        round(sum(amount),2) as balance 
from homebank_transactions 
group by name 
order by balance desc;
```