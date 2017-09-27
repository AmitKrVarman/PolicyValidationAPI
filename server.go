package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PolicyValidationAPI/model"
	"github.com/jinzhu/gorm"
	"github.com/julienschmidt/httprouter"
	"log"
)


// RegisterRouter registers a router onto the Server.
func (s *Server) RegisterRouter(router *httprouter.Router) {
	router.GET("/ping", s.ping)

	router.POST("/addPolicy", s.createPolicy)
	router.GET("/getPolicies", s.getPolicies)
	router.GET("/policy/:policyID", s.getPolicy)
}


// Server is an http server that handles REST requests.
type Server struct {
	db *gorm.DB
}

// NewServer creates a new instance of a Server.
func GetServer(db *gorm.DB) *Server {
	return &Server{db: db}
}


func (s *Server) ping(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	writeTextResult(w, "ping successful")
}

//returns all policies in JSON format
func (s *Server) getPolicies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var policies []model.Policy
	if err := s.db.Find(&policies).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, policies)
	}
}

//returns the policy specified by ID in JSON format
func (s *Server) getPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	log.Printf("getPolicy ..."+ps.ByName("policyID"))

	var policy model.Policy
	var person model.Person
	var address model.Address
	var personId int
	var addressId int

	//fetch policy details
	if err := s.db.Find(&policy, ps.ByName("policyID")).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		personId = policy.PersonID
	}

	//fetch person details
	if err := s.db.Find(&person, personId ).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		addressId = person.AddressID
	}

	//fetch address details
	if err := s.db.Find(&address, addressId ).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		person.Address = address
		policy.Person = person
	}

	writeJSONResult(w, policy)
}

//to create sample data for testing use createPolicy method
func (s *Server) createPolicy(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var policy model.Policy
	if err := json.NewDecoder(r.Body).Decode(&policy); err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
		return
	}

	if err := s.db.Create(&policy).Error; err != nil {
		http.Error(w, err.Error(), errToStatusCode(err))
	} else {
		writeJSONResult(w, policy)
	}
}

func writeTextResult(w http.ResponseWriter, res string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, res)
}

func writeJSONResult(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(res); err != nil {
		panic(err)
	}
}

func writeMissingParamError(w http.ResponseWriter, paramName string) {
	http.Error(w, fmt.Sprintf("missing query param %q", paramName), http.StatusBadRequest)
}

func errToStatusCode(err error) int {
	switch err {
	case gorm.ErrRecordNotFound:
		return http.StatusNotFound
	default:
		return http.StatusInternalServerError
	}
}




