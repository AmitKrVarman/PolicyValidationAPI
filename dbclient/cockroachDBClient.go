package dbclient

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"github.com/AmitKrVarman/PolicyValidationAPI/model"
	"flag"
	"log"
	"strconv"
)

var (
	DBAddress = flag.String("addr",
		"postgresql://root@localhost:26257/POLICYDB?sslmode=disable",
		"the address of the database")
)

func SetupDB(addr string) *gorm.DB {

	log.Printf("Setting DB ...")

	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	// Migrate the schema
	log.Printf("Migrating Schema ...")

	if db.HasTable(&model.Policy{}) { //checking one of the table
		db.AutoMigrate(&model.Policy{}, &model.Person{}, &model.Address{})
	}


	return db
}


func SeedPolicyData(addr string) {
	log.Printf("Seeding data to DB ...")

	db, err := gorm.Open("postgres", addr)
	if err != nil {
		panic(fmt.Sprintf("failed to connect to database: %v", err))
	}

	for i:=0 ; i<10 ; i++ {

		address := model.Address{
					ID:100+i,
					Address1:"15"+strconv.Itoa(i),
					Address2:"Bromegroves Street",
					Address3:"Birmingham",
					Address4:"B5"+strconv.Itoa(i)+"AE",
					}

		person := model.Person{
					ID:1000+i,
					PersonName:"Amit"+strconv.Itoa(i),
					AddressID:100+i,
					Address:address,
				}

		policy := model.Policy{
					ID : 10+i,
					Person:person,
					PersonID: 1000+i,
					Premium:float64(102.23)+float64(i),
					}

		if db.NewRecord(policy) {
			db.Create(&policy)
		} else {
			log.Printf("Policy already exist in table :  ... ID -> 10"+strconv.Itoa(i))
		}
	}

}