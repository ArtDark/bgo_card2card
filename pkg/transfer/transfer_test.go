package transfer

import (
	"github.com/ArtDark/bgo_methods/pkg/card"
	"testing"
)

func TestService_Card2Card(t *testing.T) {
	type fields struct {
		CardSvc       *card.Service
		Commission    float64
		CommissionMin int64
	}
	type args struct {
		from   string
		to     string
		amount int
	}
	tests := []struct {
		name      string
		fields    fields
		args      args
		wantTotal int
		wantOk    bool
	}{
		{
			name: "Карта своего банка -> Карта своего банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "5106 2163 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    true,
		},
		{
			name: "Карта своего банка -> Карта своего банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  34_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "5106 2163 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    false,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег достаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "4578 8942 5433 4329",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    true,
		},
		{
			name: "Карта своего банка -> Карта чужого банка (денег недостаточно)",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  34_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2158 3920 4837",
				to:     "4578 8942 5433 4329",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    false,
		},
		{
			name: "Карта чужого банка -> Карта своего банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "4578 8942 5433 4329",
				to:     "5106 2163 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    true,
		},
		{
			name: "Карта чужого банка -> Карта чужого банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "4578 8942 5433 4329",
				to:     "4534 5963 6456 3456",
				amount: 1000_00,
			},
			wantTotal: 1005_00,
			wantOk:    true,
		},
		{
			name: "Неправильная карта своего банка -> Карта чужого банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "5106 2742 5433 4321",
				to:     "4578 8942 5433 4329",
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    false,
		},
		{
			name: "Карта чужого банка -> Неправильная карта своего банка",
			fields: fields{
				CardSvc: &card.Service{
					BankName: "YourBank",
					Cards: []*card.Card{
						{
							Id: "0001",
							Owner: card.Owner{
								FirstName: "Petr",
								LastName:  "Petrov",
							},
							Issuer:   "Visa",
							Balance:  43534_34,
							Currency: "RUR",
							Number:   "5106 2158 3920 4837",
							Icon:     "http://...",
						},
						{
							Id: "0002",
							Owner: card.Owner{
								FirstName: "Ivan",
								LastName:  "Ivanov",
							},
							Issuer:   "MasterCard",
							Balance:  543534_23,
							Currency: "RUR",
							Number:   "5106 2163 6456 3456",
							Icon:     "http://...",
						},
					},
				},
				Commission:    0.5,
				CommissionMin: 10,
			},
			args: args{
				from:   "4578 8942 5433 4329",
				to:     "5106 2163 6856 3456",
				amount: 1000_00,
			},
			wantTotal: 1000_00,
			wantOk:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				CardSvc:       tt.fields.CardSvc,
				Commission:    tt.fields.Commission,
				CommissionMin: tt.fields.CommissionMin,
			}
			gotTotal, gotOk := s.Card2Card(tt.args.from, tt.args.to, tt.args.amount)
			if gotTotal != tt.wantTotal {
				t.Errorf("Card2Card() gotTotal = %v, want %v", gotTotal, tt.wantTotal)
			}
			if gotOk != tt.wantOk {
				t.Errorf("Card2Card() gotOk = %v, want %v", gotOk, tt.wantOk)
			}
		})
	}
}
