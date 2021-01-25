// Пакет для работы с транзакциями между картами
package transfer

import (
	"github.com/ArtDark/bgo_methods/pkg/card"
)

//Структура сервиса
type Service struct {
	CardSvc       *card.Service
	Commission    float64
	CommissionMin int64
}

// Функция перевода с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (total int, status bool) {

	commission := float64(amount) * s.Commission / 100.0 //Расчет комиссии
	total = amount + int(commission)                     // Расчет суммы перевода с комиссией

	fromCard, errFromCard := s.CardSvc.FindCard(from) // Поиск отправителя среди своих карт
	toCard, errToCard := s.CardSvc.FindCard(to)       // Поиск получителя среди своих карт

	if !errFromCard && !errToCard { // Если получатель и отправитель не найден среди своих карт
		return total, true
	}

	if !errFromCard { // Если отправитель не найден среди своих карт
		toCard.Balance += amount
		return total, true
	}

	if amount > fromCard.Balance { // Если баланс отправителя меньше суммы
		return total, false
	}

	if !errToCard { // Если если получатель  не найден среди своих карт
		fromCard.Balance -= total
		return total, true
	}

	fromCard.Balance -= total
	toCard.Balance += amount

	return total, true

}
