package card

type Owner struct {
	FirstName string
	LastName  string
}

type Card struct {
	Id CardId
	Owner
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}

type Service struct {
	BankName string
	Cards    []*Card
}

type CardId int64

func NewService(bankName string) *Service {
	return &Service{BankName: bankName}
}

func (s *Service) IssueCard(id CardId, fistName, lastName, issuer, currency string) *Card {
	card := &Card{
		Id: id,
		Owner: Owner{
			FirstName: fistName,
			LastName:  lastName,
		},
		Issuer:   issuer,
		Balance:  0,
		Currency: currency,
		Number:   "0001",
		Icon:     "https://.../logo.png",
	}
	s.Cards = append(s.Cards, card)
	return card
}
