package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go-boilerplate/console"
)

func main() {
	godotenv.Load()
	console.Execute()
}
