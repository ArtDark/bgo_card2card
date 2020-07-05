package main

import (
	"fmt"
	"github.com/ArtDark/bgo_card2card/pkg/card"
)

func main() {

	svc := card.NewService("art")
	fmt.Println(svc)

	visa := svc.IssueCard("Artem", "Balusov", "Visa", "RUR")

	fmt.Println(visa)

}
