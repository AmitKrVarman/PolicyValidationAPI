# PolicyValidationAPI

This API provides Rest Services for Policy ID Validation

> USE /addPolicies to add sample data
> USE /policy/policyID to retrieve Policy Data

# Pre-requisite
1. Setup Cockroch DB
2. Overide DB parameters by setting up Environment Variables (API will default if not fount)

	Variable Name		Default Value
	
	//DB Attributes
	DB_DRIVER         string = "postgres"
	DB_HOST           string = "localhost"
	DB_PORT           string = "26257"
	DB_NAME           string = "POLICYDB"
	DB_USER           string = "dbuser"
	DB_PASSWORD       string = "Password123"
	DB_SSL_MODE       string = "disable" // disable | require
	DB_MAX_CONNECTION int    = 1
	DB_LOG_MODE       bool   = true
	
	//HTTP PORT FOR API
	RUN_PORT = "6544"

3. CREATE DATABASE and set to env variables

4. Setup Tables - run API in setup mode by passing argument "true"
e.g. - $ go run *.go true



#Installation Details - Cockroach DB

How to setup Cockroch DB locally, for test purpose below steps for non secured Cluster setup ( test purpose --insecure flag has been used, can be setup as secured)

Step 1. Start the first node

	   cockroach start --insecure \
		--host=localhost

Step 2. Add nodes to the cluster

2a) In a new terminal, add the second node:

    cockroach start --insecure \
		--store=node2 \
		--host=localhost \
		--port=26258 \
		--http-port=8081 \
		--join=localhost:26257

2b) In a new terminal, add the third node:

	cockroach start --insecure \
		--store=node3 \
		--host=localhost \
		--port=26259 \
		--http-port=8082 \
		--join=localhost:26257

Step 3. Test the cluster

	cockroach sql --insecure


# PolicyValidationAPI Description

This API provide following 2 end point to get Policy Details

1. Get particular Policy Details by passing ID

Request URL    http://<server>:<port>/addPolicies
e.g - http://localhost:6544/addPolicies

Sample payload

{
	"id": 11,
	"person": {
		"id": 1001,
		"personName": "Amit Varman",
		"address": {
			"id": 101,
			"address1": "151",
			"address2": "Bromegroves Street",
			"address3": "Birmingham",
			"postcode": "B50AE"
		}
	},
	"premium": "106.23"
}



2. Get particular Policy Details by passing ID

Request URL    http://<server>:<port>/policy/{policyID}
e.g. http://localhost:6544/policy/10

Sample response

    {"id":10,"person":{"id":1000,"personName":"Amit0","address":{"id":100,"address1":"150","address2":"Bromegroves Street","address3":"Birmingham","postcode":"B50AE"},"_":100},"_":1000,"premium":"102.23"}

3. Get ALL policies (TODO add upper limit))

Request URL  http://<server>:<port>/getPolicies
e.g. http://localhost:6544/getPolicies

Sample response

    [{"id":10,"person":{"personName":"","address":{"address1":"","address2":"","address3":"","postcode":""},"_":0},"_":1000,"premium":"102.23"},{"id":11,"person":{"personName":"","address":{"address1":"","address2":"","address3":"","postcode":""},"_":0},"_":1001,"premium":"103.23"} .. ]


