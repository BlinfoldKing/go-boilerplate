package main

import (
	"go-boilerplate/console"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	console.Execute()
}
