package main

import (
	"github.com/agnjunio/gobank/api"
	"github.com/agnjunio/gobank/database"
)

func main() {
	database.Init()
	api.Init()
}
