package main

import (
	"golang-postgre-cdc/conn"
	"golang-postgre-cdc/consumer"
)

func main() {
	conn.PostgresConnect()
	consumer.PostgresConsumer()
}
