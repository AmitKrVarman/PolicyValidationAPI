package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/AmitKrVarman/PolicyValidationAPI/dbclient"
	"github.com/julienschmidt/httprouter"

	"os"
)

const ENV_RUN_PORT = "RUN_PORT"

var RUN_PORT = "6544"

func init() {
	getEnvPort()
}

func getEnvPort() {
	port := os.Getenv(ENV_RUN_PORT)
	if len(port) > 0 {
		RUN_PORT = port
	}
}

func main() {
	log.Printf("Starting API ...")

	//Open DB Connection
	db := dbclient.GetInstance()
	defer db.Close()

	//if setup mode specified - create tables
	flag.Parse()
	setupMode := flag.Arg(0)
	if setupMode == "true" {
		log.Printf("API started in Setup Mode")
		dbclient.AutoMigrate()
	}

	//GET HTTP SERVER
	server := GetServerInstance(db)
	//initialise routes
	router := httprouter.New()
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":"+RUN_PORT, router))
}
