package main

import (
	"server/loaders/fiber"
	"server/loaders/hub"
	"server/loaders/mysql"
)

func main() {
	mysql.Init()
	hub.Init()
	fiber.Init()
}
