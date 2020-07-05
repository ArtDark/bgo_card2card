package card

type Owner struct {
	FirstName string
	LastName  string
}

type Card struct {
	Id int64
	Owner
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

type Service struct {
	BankName string
	Cards    []*Service
}

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(fistName, lastName, issuer, currency) *Card {
	card := &Card{
		Id: 1 + yota,
		Owner: Owner{
			FirstName: fistName,
			LastName:  lastName,
		},
		Issuer:   issuer,
		Balance:  0,
		Currency: currency,
		Number:   "000" + string(1+yota),
		Icon:     "https://.../logo.png",
	}

	s.Card = append(s.Cards, card)

	return card
}
