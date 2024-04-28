package serviceTypes

import "backend/data/models"

type PaymentRate struct {
	PayAmount float64        `json:"pay_amount"`
	PayType   models.PayType `json:"pay_type"`
}
