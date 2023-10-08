# Sample Core Banking

## Description

API to handling all banking activity, such as:
- registered user
- deposit
- withdrawing
- check balance
- get account mutation

## Installation
### Migrate database
- build base docker image
```
make build-base-image
```
- run docker container
```
docker compose up -d
```
- run db migration
```
make migrate
```
- rollback db migration
```
make rollback
```
- start consumer service
```
make start-dev-consumer
```