package main

import (
	"fmt"
	"github.com/ArtDark/bgo_card2card/pkg/card"
	"github.com/ArtDark/bgo_card2card/pkg/transfer"
)

func main() {

	svc := card.NewService("YourBank")

	visa := svc.IssueCard(
		"0001",
		"Artem",
		"Balusov",
		"Visa",
		12345_67,
		"RUR",
		"1233_2342_2342_4322",
	)

	master := svc.IssueCard(
		"0002",
		"Ivan",
		"Ivanov",
		"MasterCard",
		98765_43,
		"RUR",
		"3242_3242_4322_2342",
	)
	fmt.Println(svc)
	fmt.Println(visa)
	fmt.Println(master)

	transferSrv := transfer.NewService(svc, 0.5, 10)

	total, status := transferSrv.Card2Card("1233_2342_2342_4322", "3242_3242_4322_2342", 5000_00)

	fmt.Println(total, status)
	fmt.Println(visa)
	fmt.Println(master)

}
