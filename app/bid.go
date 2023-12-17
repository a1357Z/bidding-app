package app

type Bid struct{
	Bidder *Buyer
	Amount int
	Status string // ACTIVE, WITHDRAWN
}