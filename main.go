package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/hidayatullahap/blog-endpoint/logging"
	"github.com/hidayatullahap/blog-endpoint/routes"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	logging.Initialize(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	logging.Trace.Log("Loggers initialized")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	routes.ApiRouter().Run(":3000")
	logging.Trace.Log("Listening on 3000")
}
