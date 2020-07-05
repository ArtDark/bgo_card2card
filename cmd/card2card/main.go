package main

import (
	"fmt"
	"github.com/ArtDark/bgo_card2card/pkg/card"
)

func main() {

	svc := card.NewService("art")
	fmt.Println(svc)

	visa := svc.IssueCard(0001, "Artem", "Balusov", "Visa", "RUR")
	master := svc.IssueCard(0002, "Ivan", "Ivanov", "MasterCard", "RUR")

	fmt.Println(visa)
	fmt.Println(master)

}
