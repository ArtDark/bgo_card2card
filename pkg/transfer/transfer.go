// Пакет для работы с транзакциями между картами
package transfer

import (
	"errors"
	"fmt"
	"github.com/ArtDark/bgo_card2card/pkg/card"
	"strconv"
	"strings"
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

// Футкция проверки номера карты
func IsValid(n string) (int, error) {
	n = strings.ReplaceAll(n, " ", "")
	sls := strings.Split(n, "")
	slsInt := [16]int{}

	for i, j := range sls {
		var err interface{}
		slsInt[i], err = strconv.Atoi(j)
		if err != nil {
			return fmt.Println(err)
		}
	}

	for _, num := range slsInt {

	}

	return 0, nil
}

// Функция перевода с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (int, error) {

	commission := float64(amount) * s.Commission / 100.0 //Расчет комиссии
	total := amount + int(commission)                    // Расчет суммы перевода с комиссией

	toCard, err := s.CardSvc.Card(to)
	if err != nil {
		toCard.Balance += amount
		return total, nil

	}
	fromCard, err := s.CardSvc.Card(from) // Поиск карты отправителя

	// Поиск карты получателя
	if fromCard == nil && toCard == nil {
		return total, nil
	}

	if toCard != nil && fromCard.Balance >= amount {
		fromCard.Balance -= int(float64(amount) + commission)
		return total, nil
	}

	if amount > fromCard.Balance { // Если баланс меньше суммы
		return amount, ErrNotEnoughMoney
	}

	fromCard.Balance -= total
	toCard.Balance += amount

	return total, nil

}
