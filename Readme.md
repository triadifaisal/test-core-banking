# Sample Core Banking

## Description

API to store user and transaction bank account

## Installation
### Migrate database
- create secret file db_password.txt and define the db password
- run docker container
```
docker compose up -d
```
- run db migration
```
make migrate url="postgres://{user_db}:{password_db}@{host_db}:{port_db}/{db_name}?sslmode=disable"
```