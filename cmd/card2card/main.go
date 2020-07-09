package main

import (
	"fmt"
	"github.com/ArtDark/bgo_card2card/pkg/card"
	"github.com/ArtDark/bgo_card2card/pkg/transfer"
)

func main() {

	yourBank := card.NewService("Your Bank")
	myBank := card.NewService("Artyom Bank")

	yourBank.IssueCard(0001, "Artyom", "Balusov", "Visa", "RUR")
	yourBank.IssueCard(0002, "Ivan", "Ivanov", "MasterCard", "RUR")
	myBank.IssueCard(0001, "Peter", "Petrov", "Visa", "EUR")
	myBank.IssueCard(0002, "Alexander", "Pushkin", "MasterCard", "RUR")

	trans := transfer.NewService(yourBank, 1, 30)

	fmt.Println(yourBank)
	fmt.Println(myBank)
	fmt.Println(trans)

}
