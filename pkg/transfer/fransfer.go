package transfer

import "github.com/ArtDark/bgo_card2card/pkg/card"

type Service struct {
	CardSvc       *card.Service
	Commission    float64 //поля для хранения:
	CommissionMin int64   // комиссий в процентах и минимума в рублях*
}

func NewService(cardSvc *card.Service, commission float64, commissionMin int64) *Service {
	return &Service{CardSvc: cardSvc, Commission: commission, CommissionMin: commissionMin}
}

func (s *Service) Card2Card(from, to string, amount int) (total int, ok bool) {
	var fromBalance *int
	var toBalance *int

	commission := float64(amount) * s.Commission / 100.0

	for _, cardNum := range s.CardSvc.Cards {
		if from == cardNum.Number {
			fromBalance = &cardNum.Balance
		} else if to == cardNum.Number {
			toBalance = &cardNum.Balance
		} else {
			return int(float64(amount) + commission), true
		}
	}

	if *fromBalance >= amount {
		*fromBalance -= int(float64(amount) + commission)
		*toBalance += amount
	} else {
		return int(float64(amount) + commission), false
	}
	return int(float64(amount) + commission), true
}
