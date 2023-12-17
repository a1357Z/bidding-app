package app

import (
	"errors"
	"fmt"
)

type Auction struct{
	creater *Seller
	lowestBid int
	highestBid int
	name string
	bids []*Bid
	status string //CLOSED
}

func createAuction (seller *Seller, low int, high int, name string) (Auction, error) {
	// can add validations
	if seller.Name == "" || low < 0 || high < 0 || (low >= high) {
		fmt.Println("cannot create auction as seller does not exist")
		return Auction{}, errors.New("cannot create auction as seller does not exist")
	}

	
	return Auction{
		creater: seller,
		lowestBid: low,
		highestBid: high,
		name: name,
	}, nil
}

func (a *Auction) GetName () string {
	return a.name
}

func (a *Auction) CreateBid (bidder *Buyer, amount int) {

	if bidder.Name == "" || amount < a.lowestBid || amount > a.highestBid{
		fmt.Println("cannot create bid")
		return
	}

	// bid is exclusive to auction 
	bid := Bid{
		Bidder: bidder,
		Amount: amount,
		Status: "ACTIVE",
	}

	a.bids = append(a.bids, &bid)
	fmt.Printf("Created bid by: %s, of amount: %d\n", bid.Bidder.Name, bid.Amount)
}

func (a *Auction) findBid (bidder *Buyer) (*Bid, error) {
	
	for _, bid := range a.bids {
		if bid.Bidder.Name == bidder.Name{
			return bid, nil
		}
	}

	fmt.Println("Bid does not exist")
	return &Bid{}, errors.New("Bid does not exist")
}

func (a *Auction) listBids () {
	fmt.Printf("Listing bids for auction: %v\n", a.name)

	for _, bid := range a.bids{
		fmt.Printf("Bidder: %s, Amount: %d\n", bid.Bidder.Name, bid.Amount)
	}
}

func (a *Auction) UpdateBid (bidder *Buyer, amount int) {
	if bidder.Name == "" || amount < a.lowestBid || amount > a.highestBid{
		fmt.Println("cannot update bid")
		return
	}

	// bid is exclusive to auction 
	bid, err := a.findBid(bidder)
	if err != nil{
		return
	}

	// can create a method to update amount to prevent directly updating the bid
	bid.Amount = amount
	fmt.Printf("Updated bid by %s, amount: %d for auction: %s\n", bid.Bidder.Name, bid.Amount, a.name)
}

func (a *Auction) WithdrawBid (bidder *Buyer) {
	if bidder.Name == "" {
		fmt.Println("cannot withdraw bid")
		return
	}

	// bid is exclusive to auction 
	bid, err := a.findBid(bidder)
	if err != nil{
		fmt.Println("cannot withdraw bid")
		return
	}

	bid.Status = "WITHDRAWN"
	fmt.Printf("Withdrawn bid by %s, for auction: %s\n", bid.Bidder.Name, a.name)
}

func (a *Auction) findWinningBid () (*Bid, error) {
	winningBid := Bid{}
	for _, bid := range a.bids{
		if bid.Status != "WITHDRAWN" && bid.Amount <= a.highestBid && bid.Amount >= a.lowestBid && bid.Amount > winningBid.Amount {
			winningBid = *bid
		}
	}

	if winningBid.Amount == 0 {
		fmt.Println("no winning bid found")
		return &Bid{}, errors.New("no winning bid found")
	}

	return &winningBid, nil
}

func (a *Auction) Close () {
	
	bid, err := a.findWinningBid()

	a.status = "CLOSED"
	fmt.Printf("Closed auction: %s\n", a.name)

	if err == nil{
		fmt.Printf("Found winner: %s, with winning bid of %d\n", bid.Bidder.Name, bid.Amount)
	}
	
}




