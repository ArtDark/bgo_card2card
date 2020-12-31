// Пакет управления банковскими картами
package card

import (
	"errors"
	"strings"
)

// Описание банковской карты
type Card struct {
	Id       cardId // Идентификатор карты в системе банка
	Owner           // Владелец карты
	Issuer   string // Платежная система
	Balance  int    // Баланс карты
	Currency string // Валюта
	Number   string // Номер карты в платежной системе
	Icon     string // Иконка платежной системы
}

// Идентификатор банковской карты
type cardId string

// Инициалы владельца банковской карты
type Owner struct {
	FirstName string // Имя владельца карты
	LastName  string // Фамилия владельца карты
}

// Сервис банка
type Service struct {
	BankName string
	Cards    []*Card
}

// Конструктор сервиса
func New(bankName string) *Service {
	return &Service{BankName: bankName}
}

// Метод создания экземпляра банковской карты
func (s *Service) CardIssue(
	fistName,
	lastName,
	issuer string,
	currency string,
) *Card {
	var card = &Card{
		Owner: Owner{
			FirstName: fistName,
			LastName:  lastName,
		},
		Issuer:   issuer,
		Currency: currency,
		Icon:     "https://.../logo.png",
	}
	s.Cards = append(s.Cards, card)
	return card
}

var ErrCardNotFound = errors.New("card not found")

const prefix = "5106 21" //Первые 6 цифр нашего банка

// Метод поиска банковской карты по номеру платежной системы
func (s *Service) Card(number string) (*Card, error) {

	for _, с := range s.Cards {
		if strings.HasPrefix(с.Number, prefix) == true {
			return с, nil
		}
	}
	return nil, ErrCardNotFound
}
