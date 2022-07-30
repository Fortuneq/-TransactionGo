# -TransactionGo
Механизм транзакций на GO

### Setup infrastructure

- Run db migration:

    ```bash
    make migrate_up
    ```

### How to generate code

- Generate SQL CRUD with sqlc:

    ```bash
    make sqlc
    ```

- Generate DB mock with gomock:

    ```bash
    make mock
    ```

### How to run

- Run server:

    ```bash
    make server
    ```

- Run test:

    ```bash
    make test
    ```
