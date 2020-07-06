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

//func (s *Service) Card2Card(from, to string, amount int) (total int, ok bool) {
//
//}
