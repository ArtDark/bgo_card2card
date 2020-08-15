// Пакет для работы с транзакциями между картами
package transfer

import (
	"errors"
	"github.com/ArtDark/bgo_card2card/pkg/card"
)

//Структура сервиса
type Service struct {
	CardSvc       *card.Service
	Commission    float64
	CommissionMin int64
}

func NewService(cardSvc *card.Service, commission float64, commissionMin int64) *Service {
	return &Service{CardSvc: cardSvc, Commission: commission, CommissionMin: commissionMin}
}

var (
	ErrNotEnoughMoney = errors.New("not enough money")
)

// Функция перевода с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (int, error) {

	commission := float64(amount) * s.Commission / 100.0 //Расчет комиссии
	total := amount + int(commission)                    // Расчет суммы перевода с комиссией

	fromCard := s.CardSvc.Card(from) // Поиск карты отправителя
	toCard := s.CardSvc.Card(to)     // Поиск карты получателя

	if fromCard == nil && toCard == nil { // Если нет наших карт
		return amount, nil
	}

	if fromCard == nil {
		toCard.Balance += amount
		return total, nil
	}

	if toCard == nil && fromCard.Balance >= amount {
		fromCard.Balance -= int(float64(amount) + commission)
		return total, nil
	}

	if amount > fromCard.Balance { // Если баланс меньше суммы
		return amount, ErrNotEnoughMoney
	}

	fromCard.Balance -= int(float64(amount) + commission)
	toCard.Balance += amount

	return int(float64(amount) + commission), nil

}
