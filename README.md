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

### Запуск кода
![Запуск кода ](screenchots/Screenshotfrom2022-07-3101-10-47.png?raw=true "Запуск кода")
### Видим рандомно сгенерированных пользователей
![База данных ](screenchots/Screenshotfrom2022-07-3101-11-43.png?raw=true "Видим рандомно сгенерированных пользователей")
### Попытка трансфера с недостатком средств
![Недостаток средств](screenchots/Screenshotfrom2022-07-3101-12-33.png?raw=true "Попытка трансфера с недостатком средств")
### Обычный трансфер, по клиенту строится json вывод о балансе двух акаунтов и для разработчика также можно увидеть время когда произошёл вход в транзакцию и id трансфера 
![Обычный трансфер](screenchots/Screenshotfrom2022-07-3101-14-43.png?raw=true "Обычный трансфер, по клиенту строится json вывод о балансе двух акаунтов и для разработчика также можно увидеть время когда произошёл вход в транзакцию и id трансфера ")

## Система очередей создана на Redis при создании каждого трансфера создаётся очередь которая возвращает json  
## Сохранение данных при вылете сервера нам обеспечивает журнал транзакций postgresql который находится на диске
