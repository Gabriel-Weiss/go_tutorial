package models

type Cart struct {
	ID       string
	Quantity int
	Buyer    string
}

var Products = []Cart{
	{
		ID:       "Nero",
		Quantity: 2,
		Buyer:    "user1",
	},
	{
		ID:       "Addidas",
		Quantity: 1,
		Buyer:    "user1",
	},
	{
		ID:       "Fossil",
		Quantity: 4,
		Buyer:    "user1",
	},
}
