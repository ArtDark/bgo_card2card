// Пакет для работы с транзакциями между картами
package transfer

import (
	"github.com/ArtDark/bgo_methods/pkg/card"
	"strconv"
	"strings"
)

//Структура сервиса
type Service struct {
	CardSvc       *card.Service
	Commission    float64
	CommissionMin int64
}

// Функция проверки номера карты
func IsValid(n string) bool {
	n = strings.ReplaceAll(n, " ", "") // Удаление пробелов из строки
	if len(n) != 16 {
		return false
	}
	sls := strings.Split(n, "") // Создание слайса из строки
	slsInt := [16]int{}         // Создание слайса типа int

	// Преобразование значение string -> int, запись в слайс int
	for i, j := range sls {
		var err interface{}
		slsInt[i], err = strconv.Atoi(j)
		if err != nil {
			return false
		}
	}
	// Операция над каждым нечетным числом с последующим изменением в слайсе slsInt
	for i := 0; i < len(slsInt); i += 2 {
		num := slsInt[i] * 2

		if num > 9 {
			num -= 9
		}

		slsInt[i] = num
	}

	sum := 0 // Контрольная сумма

	// Сумма всех чисел в слайсе
	for _, i := range slsInt {
		sum += i
	}

	// Проверка на кратность 10
	if sum%10 == 0 {
		return true
	}

	return false
}

// Функция перевода с карты на карту
func (s *Service) Card2Card(from, to string, amount int) (total int, status bool) {

	if !IsValid(from) || !IsValid(to) {
		return amount, false
	}

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
