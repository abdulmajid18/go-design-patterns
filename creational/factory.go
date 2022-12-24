package creational

import "errors"

type PaymentMethod interface {
	Pay(amount float32) string
}

const (
	Cash      = 1
	DebitCard = 2
)

func GetPaymentMethod(m int) (PaymentMethod, error) {
	return nil, errors.New("not implementedy yet")
}

type CashPM struct{}
type DebitCardPM struct{}

func (c *CashPM) Pay(amount float32) string {
	return ""
}

func (c *DebitCardPM) Pay(amount float32) string {
	return ""
}
