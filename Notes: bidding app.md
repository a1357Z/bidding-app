Notes: bidding app

Problem Statement:

A company SuperBidder has hosted an Online Auction System where any user(seller) can sell an object through an auction. Each auction has a lowest bid limit and the highest bid limit. Any registered user(buyer) can participate in an auction and bid on the product. Buyers can update this bid amount or withdraw from an auction until the auction is completed. When the auction closes, show the winning bid using the highest unique bid. The program should take as input two or more auctions and a set of users participating in these. Multiple auctions can happen simultaneously. A new requirement came in where SuperBidder has requested that they want to decide the winner based on the lowest unique bid as well. Accommodate this change as well. For SDE 3s only Buyers have a limited total budget in the beginning. If the buyer wins the auction, his budget is reduced by the bid amount. Any new bid exceeding the bid amount should be rejected. Buyers should be able to increase the budget at any point as well. Highest Unique bid definition For a set of users A, B, C, D, E participating in auction A1 A bids 50, B bids 90, C bids 100, D bids 90, E bids 70, F bids 100 Here 70 is the highest unique bid, therefore E is the winner. If there is no highest unique bid by the end of the auction, there is no winner for the auction. Functional Requirements • Add buyer • Add seller • Create auction • Create/Update bid • Withdraw bid • Close auction and return winning bid and the winner name • Update budget (Only for SDE 3s) Bonus • Upgrade the buyer to a preferred buyer if he has participated in more than 2 auctions. And for choosing a winner, whenever there is a tie on the winning bid, preference should be given to the preferred buyer and if it’s tied between multiple preferred buyers, fallback to the next highest bid. • The preferred buyer is across sellers on the platform.


Sample Input Used:

ADD_SELLER SHUBHAM
ADD_SELLER AKASH

CREATE_AUCTION BISCUITS 10 25 SHUBHAM
CREATE_AUCTION NAMKEEN 10 25 SHUBHAM
CREATE_AUCTION DARU 10 25 SHUBHAM
CREATE_AUCTION BEER 10 25 SHUBHAM

ADD_BUYER RACHIT
ADD_BUYER AMIT 
ADD_BUYER ARNAV 

CREATE_BID RACHIT BISCUITS 15
CREATE_BID RACHIT DARU 15
CREATE_BID AMIT BEER 15
CREATE_BID ARNAV BEER 15
CREATE_BID ARNAV NAMKEEN 14

UPDATE_BID AMIT BEER 19
UPDATE_BID AMIT BISCUITS 19

WITHDRAW_BID RACHIT BISCUITS

CLOSE_AUCTION BISCUITS
CLOSE_AUCTION NAMKEEN
CLOSE_AUCTION DARU
CLOSE_AUCTION BEER

----------------------------------------------------------

Functional Requirements :
 • Add buyer 
 • Add seller 
 • Create auction 
 • Create/Update bid 
 • Withdraw bid 
 • Close auction and return winning bid and the winner name 
 
 <!-- for later -->
 • Update budget (Only for SDE 3s) 
 Bonus • Upgrade the buyer to a preferred buyer if he has participated in more than 2 auctions. And for choosing a winner, whenever there is a tie on the winning bid, preference should be given to the preferred buyer and if it’s tied between multiple preferred buyers, fallback to the next highest bid. • The preferred buyer is across sellers on the platform.

Time distribution:
get clarification: 15 minutes
15 minutes per requirement

Queries:
can a user be both buyer and seller ? Assumption No.

Assumptions:
    Auction name is unique and can be treated as the id


Entities:

User:
    Name
    ID


Buyer --> composed of user
    List of Auctions
    Name
    ID
    Methods:
        Create bid
        update bid
        withdraw bid

Seller --> composed of user
    List of Auctions
    Name
    ID
    Methods:
        Create auction

Auction -->
    Buyers: List
    Seller
    Methods:
        GetWinner

System:
    Buyers: List
    Sellers: List
    list of Auctions

Issues faced while coding:
    Not passing in pointers resulted in empty values of bids list
    Issue in taking in space separated input
    Could not complete in 2.5 hrs in the first time.