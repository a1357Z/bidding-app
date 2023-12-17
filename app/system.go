package app

import (
	"errors"
	"fmt"
	// "strings"
)

type System struct{
	sellers []*Seller
	buyers []*Buyer
	auctions []*Auction
}

func GetSystem() System {
	return System{
		sellers: make([]*Seller, 0),
		buyers: make([]*Buyer, 0),
		auctions: make([]*Auction, 0),
	}
}

// how to handle invalid input ?
func (s *System) ProcessInput () bool {
	var input string

	_, err := fmt.Scan(&input)

	if err != nil{
		fmt.Printf("Error in scanning input: %s\n", err.Error())
	}

	if input == "CLOSE" {
		return true
	}
	
	if input == "ADD_SELLER" {
		s.addSeller()
		return false
	}

	if input == "ADD_BUYER" {
		s.addBuyer()
		return false
	}

	if input == "CREATE_AUCTION" {
		s.createAuction()
		return false
	}

	if input == "CREATE_BID" {
		s.createBid()
		return false
	}

	if input == "UPDATE_BID" {
		s.updateBid()
		return false
	}

	if input == "WITHDRAW_BID" {
		s.withdrawBid()
		return false
	}

	if input == "CLOSE_AUCTION" {
		s.closeAuction()
		return false
	}

	fmt.Println("invalid input")
	return false
	
}

func (s *System) listAllAuctionsBids () {
	for _, a := range s.auctions{
		a.listBids()
	}
}

// for seller
func (s *System) findSeller (name string) *Seller {
	for _, seller := range s.sellers {
		if seller.Name == name {
			return seller
		}
	}
	return &Seller{}
}

func (s *System) sellerExists (name string) bool {
	seller := s.findSeller(name)
	return seller.Name == name
}

func (s *System) addSeller () {
	// skip input validation for later

	var name string
	_, err := fmt.Scan(&name)

	if err != nil{
		fmt.Printf("Error in scanning addSeller input: %s\n", err.Error())
		return
	}

	// validate if a seller already exists with same name?
	if s.sellerExists(name) {
		fmt.Println("Seller already exists")
		return
	}

	// create seller
	seller := Seller{
		Name: name,
	}

	s.sellers = append(s.sellers, &seller)

	fmt.Printf("Added Seller: %v\n", seller)
}

// for buyer
func (s *System) findBuyer (name string) *Buyer {
	for _, buyer := range s.buyers {
		if buyer.Name == name {
			return buyer
		}
	}
	return &Buyer{}
}

func (s *System) buyerExists (name string) bool {
	for _, buyer := range s.buyers {
		if buyer.Name == name {
			return true
		}
	}

	return false
}

func (s *System) addBuyer () {
	// skip input validation for later

	var name string
	_, err := fmt.Scan(&name)

	if err != nil{
		fmt.Printf("Error in scanning addBuyer input: %s\n", err.Error())
		return
	}

	// validate if a seller already exists with same name?
	if s.buyerExists(name) {
		fmt.Println("Buyer already exists")
		return
	}

	// create buyer
	// TODO: move to buyer.go and encapsulate the buyer struct
	buyer := Buyer{
		Name: name,
	}

	s.buyers = append(s.buyers, &buyer)

	fmt.Printf("Added buyer: %v\n", buyer)
}

func (s *System) createAuction () {
	var name, seller string
	_, err := fmt.Scan(&name)

	if err != nil{
		fmt.Printf("Error in scanning createAuction name: %s\n", err.Error())
		return
	}

	var low, high int

	_, err = fmt.Scan(&low)
	if err != nil{
		fmt.Printf("Error in scanning createAuction lowbid: %s\n", err.Error())
		return
	}

	_, err = fmt.Scan(&high)
	if err != nil{
		fmt.Printf("Error in scanning createAuction highbid: %s\n", err.Error())
		return
	}

	_, err = fmt.Scan(&seller)
	if err != nil{
		fmt.Printf("Error in scanning createAuction seller: %s\n", err.Error())
		return
	}

	// find the seller first
	creator := s.findSeller(seller)
	_, err = s.findAuction(name)
	
	if err == nil {
		fmt.Println("Auction already exists")
		return
	}

	auction, err := createAuction(creator, low, high, name)

	if err != nil{
		return
	}

	fmt.Printf("Created Auction: %v\n", auction)
	s.auctions = append(s.auctions, &auction)
}

func (s *System) findAuction (name string) (*Auction, error) {
	for _, auction := range s.auctions {
		if auction.GetName() == name {
			return auction, nil
		}
	}
	return &Auction{}, errors.New("auction not found")
}

func systemScan (input *interface{}, inputName string, inputMethod string) {
	_, err := fmt.Scan(&input)

	if err != nil{
		fmt.Printf("Error in scanning %s %s: %s\n", inputMethod, inputName, err.Error())
		return
	}
}

func (s *System) createBid () {
	var bidderName, auctionName string

	_, err := fmt.Scan(&bidderName)

	if err != nil{
		fmt.Printf("Error in scanning createBid bidderName: %s\n", err.Error())
		return
	}

	_, err = fmt.Scan(&auctionName)

	if err != nil{
		fmt.Printf("Error in scanning createBid auctionName: %s\n", err.Error())
		return
	}

	var bidAmount int
	_, err = fmt.Scan(&bidAmount)

	if err != nil{
		fmt.Printf("Error in scanning createBid bidAmount: %s\n", err.Error())
		return
	}

	auction, err := s.findAuction(auctionName)
	
	if err != nil {
		fmt.Println("Auction not found")
	}

	bidder := s.findBuyer(bidderName)

	auction.CreateBid(bidder, bidAmount)
}

// TODO: move the common logic of create and update bid here
func (s *System) inputBid() () {}

func (s *System) updateBid () {
	var bidderName, auctionName string

	_, err := fmt.Scan(&bidderName)

	if err != nil{
		fmt.Printf("Error in scanning updateBid bidderName: %s\n", err.Error())
		return
	}

	_, err = fmt.Scan(&auctionName)

	if err != nil{
		fmt.Printf("Error in scanning updateBid auctionName: %s\n", err.Error())
		return
	}

	var bidAmount int
	_, err = fmt.Scan(&bidAmount)

	if err != nil{
		fmt.Printf("Error in scanning updateBid bidAmount: %s\n", err.Error())
		return
	}

	auction, err := s.findAuction(auctionName)
	
	if err != nil {
		fmt.Println("Auction not found")
	}

	bidder := s.findBuyer(bidderName)

	// s.listAllAuctionsBids()
	auction.UpdateBid(bidder, bidAmount)
}



func (s *System) withdrawBid () {
	var bidderName, auctionName string

	_, err := fmt.Scan(&bidderName)

	if err != nil{
		fmt.Printf("Error in scanning updateBid bidderName: %s\n", err.Error())
		return
	}

	_, err = fmt.Scan(&auctionName)

	if err != nil{
		fmt.Printf("Error in scanning updateBid auctionName: %s\n", err.Error())
		return
	}

	auction, err := s.findAuction(auctionName)
	
	if err != nil {
		fmt.Println("Auction not found")
	}

	bidder := s.findBuyer(bidderName)

	// s.listAllAuctionsBids()
	auction.WithdrawBid(bidder)
}

func (s *System) closeAuction () {
	var  auctionName string
	_, err := fmt.Scan(&auctionName)

	if err != nil{
		fmt.Printf("Error in scanning updateBid auctionName: %s\n", err.Error())
		return
	}

	auction, err := s.findAuction(auctionName)
	
	if err != nil {
		fmt.Println("Auction not found")
	}

	auction.Close()
}

