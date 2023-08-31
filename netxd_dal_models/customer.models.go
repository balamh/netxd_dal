package netxd_dal_models

import "time"
type NetxdCustomer struct{
	CustomerId int32 `json:"customerid" bson:"customerid"`
	FirstName string `json:"firstname" bson:"firstname"`
	LastName string `json:"lastname" bson:"lastname"`
	BankId int32 `json:"bankid" bson:"bankid"`
	Balance float64 `json:"balance" bson:"balance"`
	CreatedAt time.Time  `json:"createdat" bson:"createdat"`
	UpdatedAt  time.Time `json:"updatedat" bson:"updatedat"`
	IsActive bool `json:"isactive" bson:"isactive"`
}

type CustomerResponse struct{
	CustromerId int32 `json:"customerid" bson:"customerid"`
	CreatedAt time.Time  `json:"createdat" bson:"createdat"`
}