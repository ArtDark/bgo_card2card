// Пакет управления банковскими картами
package card

// Описание банковской карты
type Card struct {
	Id       cardId // Идентификатор карты в системе банка
	Owner           // Владелец карты
	Issuer   string // Длатежная истема
	Balance  int    // Баланс карты
	Currency string // Валюта
	Number   string // Номер карты в платежной системе
	Icon     string // Иконка платежной системы
}

// Идентификат банковской карты
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
	id cardId,
	fistName,
	lastName,
	issuer string,
	balance int,
	currency string,
	number string,
) *Card {
	var card = &Card{
		Id: id,
		Owner: Owner{
			FirstName: fistName,
			LastName:  lastName,
		},
		Issuer:   issuer,
		Balance:  balance,
		Currency: currency,
		Number:   number,
		Icon:     "https://.../logo.png",
	}
	s.Cards = append(s.Cards, card)
	return card
}
