package creational

import (
	"errors"
	"fmt"
)

const (
	Cash = 1
	Card = 2
)

type PaymentMethod interface {
	Pay(amount float32) string
}

type CashPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f paid using cash", amount)
}

type CardPM struct{}

func (c *CardPM) Pay(amount float32) string {
	return fmt.Sprintf("%.2f paid using card (new implementation )", amount)
}

// Factory of payment methods
func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPM{}, nil
	case Card:
		return &CardPM{}, nil
	default:
		return nil, errors.New("unrecognised payment method")
	}
}
