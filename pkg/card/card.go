package card

type Owner struct {
	FirstName string
	LastName  string
}

type Card struct {
	Id cardId
	Owner
	Issuer   string
	Balance  int
	Currency string
	Number   string
	Icon     string
}

type Service struct {
	BankName string
	Cards    []*Card
}

type cardId string

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(
	id cardId,
	fistName,
	lastName,
	issuer string,
	balance int,
	currency string,
	number string,
) *Card {
	var card = &Card{
		Id: id,
		Owner: Owner{
			FirstName: fistName,
			LastName:  lastName,
		},
		Issuer:   issuer,
		Balance:  balance,
		Currency: currency,
		Number:   number,
		Icon:     "https://.../logo.png",
	}
	s.Cards = append(s.Cards, card)
	return card
}
