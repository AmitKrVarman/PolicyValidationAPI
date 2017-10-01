package main

import (
	"flag"
	"log"
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/julienschmidt/httprouter"
	"github.com/AmitKrVarman/PolicyValidationAPI/dbclient"

)


func main() {
	log.Printf("App Started ...")
	flag.Parse()

	db := dbclient.SetupDB(*dbclient.DBAddress)
	//dbclient.SeedPolicyData(*dbclient.DBAddress)

	defer db.Close()

	router := httprouter.New()

	server := GetServer(db)
	server.RegisterRouter(router)

	log.Fatal(http.ListenAndServe(":6543", router))
}


