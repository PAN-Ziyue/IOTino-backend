# IOTino-Backend

[![Build And Test](https://github.com/PAN-Ziyue/IOTino-backend/workflows/CI/badge.svg?event=push)](https://github.com/PAN-Ziyue/IOTino-backend/actions?workflow=CI)

The backend project for IOTino.

## Dependencies

- Go
- MySQL

## Config

Edit `config/IOTino.ini`, an example:

```
[MQTT]
MQTTAddr = localhost
MQTTPort = 1883
KeepAlive = 20


[Server]
ServerAddr = localhost
RunMode = debug
HTTPPort = 1882
READ_TIMEOUT = 60
WRITE_TIMEOUT = 60


[Database]
TYPE = mysql
USER = root
PASSWORD = 600019
HOST = 127.0.0.1:3306
TABLE = IOTino
```

## Build

```bash
$ go mod tidy # sync dependencies
$ go build IOTino # build binary
```

## Run

```bash
$ ./IOTino
```
