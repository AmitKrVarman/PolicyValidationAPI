package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/AmitKrVarman/PolicyValidationAPI/dbclient"

	"os"
	"fmt"
	"github.com/jinzhu/gorm"
)


func main() {
	log.Printf("Strating API ...")
	flag.Parse()
	setupMode:= flag.Arg(0) //read setup mode

	dbAddr := os.Getenv("DB_ADDRESS") //GET DB ADDRESS
	log.Printf("DB_ADDRESS :- "+dbAddr )


	//Open DB Connection
	db, err := gorm.Open("postgres", dbAddr)
	//Create DB
	db.Exec("CREATE DATABASE IF NOT EXISTS POLICYDB");

	if  err != nil {
		log.Fatal(err)
	}

	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	} else {
		log.Printf("DB connection establisded...")
		defer db.Close()
	}

	//one time set to create Tables if
	if setupMode == "true" {
		log.Printf("API started in Setup Mode")
		dbclient.SetupDB(db)
		dbclient.SeedPolicyData(db)
	}

	router := httprouter.New()

	server := GetServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":6543", router))
}


