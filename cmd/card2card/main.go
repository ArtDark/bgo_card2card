package main

import (
	"fmt"
	"github.com/ArtDark/bgo_methods/pkg/card"
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

}
