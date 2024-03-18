package main

import (
	"fmt"
	"os"

	"github.com/ibad69/go.fields/pkg/api"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("server starting")
	godotenv.Load("../../.env")

	uri := os.Getenv("db_uri")
	fmt.Printf("uri: %s\n", uri)
	api.Start(&api.Config{
		DbHost: uri,
		DbPort: 22,
		DbUser: "string",
		DbPw:   "string",
		DbName: "string",

		AppHost: "string",
		AppPort: os.Getenv("PORT"),
	})
}
