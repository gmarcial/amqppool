# AMQP Pool
A simple amqp channel pool made from of go client [streadway/amqp](https://github.com/streadway/amqp).

### WARNING: Immature, don't was tested them enough in production.

## Install
    go get github.com/gmarcial/amqp-pool

## Usage
- [Example](./example/main.go)

## Test
Configure an instance of rabbitmq on your machine, export the connection string how environment variable and run the tests.

Example:

    export AMQP_CONNECTION=amqp://guest:guest@127.0.0.1:5672/ 

    go test ./...