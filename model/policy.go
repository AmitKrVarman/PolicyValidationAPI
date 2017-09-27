package model

// Product is a model in the "products" table.
type Policy struct {
	ID    int     `json:"id,omitempty"`

	Person   Person `json:"person" gorm:"ForeignKey:PersonID"`
	PersonID int    `json:"_"`

	Premium float64 `json:"premium,string" gorm:"type:decimal(10,2)"`
}

// Person is a model in the "Person" table.
type Person struct {
	ID    int     `json:"id,omitempty"`
	PersonName  string `json:"personName"  gorm:"not null;"`

	Address Address `json:"address" gorm:"ForeignKey:AddressID"`
	AddressID int `json:"_"`
}


// RiskAddress is a model in the "RiskAddress" table.
type Address struct {
	ID    int     `json:"id,omitempty"`

	Address1  string `json:"address1"  gorm:"not null;"`
	Address2  string `json:"address2"`
	Address3 string  `json:"address3"`
	Address4  string `json:"postcode"  gorm:"not null;"`
}
