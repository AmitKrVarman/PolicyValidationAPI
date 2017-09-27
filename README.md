# PolicyValidationAPI

## Pre-requisite

### Cockroch DB

How to setup Cockroch DB locally, for test purpose below steps for non secured Cluster setup

Step 1. Start the first node

	   cockroach start --insecure --host=localhost

Step 2. Add nodes to the cluster

2a) In a new terminal, add the second node:

    cockroach start --insecure --store=node2 --host=localhost --port=26258 --http-port=8081 --join=localhost:26257

2b) In a new terminal, add the third node:

	cockroach start --insecure --store=node3 --host=localhost --port=26259 --http-port=8082 --join=localhost:26257

Step 3. Test the cluster

	cockroach sql --insecure

# PolicyValidationAPI Description

This API provide following 2 end point to get Policy Details

1. Get particular Policy Details by passing ID

Request URL    http://<server>:<port>/policy/{policyID}

Sample response

    {"id":10,"person":{"id":1000,"personName":"Amit0","address":{"id":100,"address1":"150","address2":"Bromegroves Street","address3":"Birmingham","postcode":"B50AE"},"_":100},"_":1000,"premium":"102.23"}

2. Get list of policies - for testing purpose 10 policy data would be seeded to Cochroch DB on startup

Request URL  http://<server>:<port>/getPolicies

Sample response

    [{"id":10,"person":{"personName":"","address":{"address1":"","address2":"","address3":"","postcode":""},"_":0},"_":1000,"premium":"102.23"},{"id":11,"person":{"personName":"","address":{"address1":"","address2":"","address3":"","postcode":""},"_":0},"_":1001,"premium":"103.23"} .. ]


### Customisation

cockroachDBClient.go : Update client code to use remote DB Cluster instead of localhost

var (
	DBAddress = flag.String("addr",
		"postgresql://root@localhost:26257/POLICYDB?sslmode=disable",
		"the address of the database")
)