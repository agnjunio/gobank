package main

import (
	"github.com/agnjunio/gobank/internal/app"
	"github.com/agnjunio/gobank/internal/common/database"
)

func main() {
	database.Init()
	app.Init()
}
