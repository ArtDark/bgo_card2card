package card

type Owner struct {
	FirstName string
	LastName  string
}

type Card struct {
	Id int64
	Owner
	Issuer   string
	Balance  int64
	Currency string
	Number   string
	Icon     string
}
