package models

import (
	"errors"
	"fmt"
	"time"
)

// PaymentRate represents the rate at which someone is paid.
type PaymentRate struct {
	PayAmount float64 // The amount to be paid per unit (e.g., per hour, per week, etc.)
	PayType   PayType // The type of payment rate (e.g., hourly, weekly, etc.)
}

// NewPaymentRate creates a new PaymentRate.
func NewPaymentRate(amount float64, payType PayType) *PaymentRate {
	return &PaymentRate{
		PayAmount: amount,
		PayType:   payType,
	}
}

// CalculateAmount calculates the total amount earned based on the rate and the provided units.
// For hourly, units represent hours.
// For weekly, units represent weeks, and similarly for monthly and yearly.
func (pr *PaymentRate) CalculateAmount(units float64) (float64, error) {
	if units < 0 {
		return 0, errors.New("units cannot be negative")
	}

	switch pr.PayType {
	case Hourly:
		return pr.PayAmount * units, nil
	case Weekly:
		return pr.PayAmount * units, nil
	case Monthly:
		return pr.PayAmount * units, nil
	case Yearly:
		return pr.PayAmount * units, nil
	default:
		return 0, fmt.Errorf("unknown pay type: %s", pr.PayType)
	}
}

func (pr *PaymentRate) CalculatePayoutForPeriod(start, end time.Time) (float64, error) {
	if end.Before(start) {
		return 0, errors.New("end time cannot be before start time")
	}

	duration := end.Sub(start)

	switch pr.PayType {
	case Hourly:
		hours := duration.Hours()
		return pr.PayAmount * hours, nil
	case Weekly:
		weeks := duration.Hours() / (7 * 24)
		return pr.PayAmount * weeks, nil
	case Monthly:
		months := duration.Hours() / (30.44 * 24)
		return pr.PayAmount * months, nil
	case Yearly:
		years := duration.Hours() / (365.24 * 24)
		return pr.PayAmount * years, nil
	default:
		return 0, fmt.Errorf("unknown pay type: %s", pr.PayType)
	}
}
