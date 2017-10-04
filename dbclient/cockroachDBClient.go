package dbclient

import (
	"github.com/jinzhu/gorm"
	"fmt"
	"flag"
	"log"
	"strconv"
	"github.com/AmitKrVarman/PolicyValidationAPI/model"
)

/*var (
	DBAddress = flag.String("addr",
		"postgresql://root@localhost:26257/POLICYDB?sslmode=disable",
		"the address of the database")
)*/

func SetupDB(db *gorm.DB) {
	// Migrate the schema
	log.Printf("Auto Migrating Schemas ...")
	db.AutoMigrate(&model.Policy{}, &model.Person{}, &model.Address{})

}


func SeedPolicyData(db *gorm.DB) {
	log.Printf("Seeding data to DB ...")

	for i:=11 ; i<20 ; i++ {

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