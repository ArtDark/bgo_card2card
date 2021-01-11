package main

import (
	"fmt"
	"github.com/ArtDark/bgo_methods/pkg/card"
	"github.com/ArtDark/bgo_methods/pkg/transfer"
)

func main() {

	svc := card.New("YourBank")

	visa := svc.CardIssue(
		"Petr",
		"Petrov",
		"Visa",
		"RUR",
	)

	master := svc.CardIssue(
		"Ivan",
		"Ivanov",
		"MasterCard",
		"RUR",
	)

	fmt.Println(svc)
	fmt.Println(visa)
	fmt.Println(master)

	s := transfer.IsValid("4  5  6  1     2  6  1  2     1  2  3  4     5  4  6  7")

	fmt.Println(s)

}
