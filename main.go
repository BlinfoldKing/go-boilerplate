package main

import (
	"go-boilerplate/config"
	"go-boilerplate/console"
	"go-boilerplate/helper"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	helper.InitLogger(config.ENV())
	console.Execute()
}
