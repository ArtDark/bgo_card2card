// Пакет для работы с транзакциями между картами
package transfer

import "github.com/ArtDark/bgo_card2card/pkg/card"

type Service struct {
	CardSvc       *card.Service
	Commission    float64
	CommissionMin int64
}

func NewService(cardSvc *card.Service, commission float64, commissionMin int64) *Service {
	return &Service{CardSvc: cardSvc, Commission: commission, CommissionMin: commissionMin}
}

func (s *Service) Card2Card(from, to string, amount int) (total int, ok bool) {

	commission := float64(amount) * s.Commission / 100.0

	fromCard := s.CardSvc.Card(from)
	toCard := s.CardSvc.Card(to)

	if fromCard == nil && toCard == nil {
		return int(float64(amount) + commission), false
	}

	if fromCard == nil {
		toCard.Balance += amount
		return int(float64(amount) + commission), true

	}

	if toCard == nil {
		fromCard.Balance -= int(float64(amount) + commission)
		return int(float64(amount) + commission), true

	}

	if fromCard.Balance >= amount {
		fromCard.Balance -= int(float64(amount) + commission)
		toCard.Balance += amount
	} else {
		return int(float64(amount) + commission), false
	}
	return int(float64(amount) + commission), true
}
