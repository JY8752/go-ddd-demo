package money

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/shopspring/decimal"
)

type Currency int

const (
	_ Currency = iota
	JPY
	USD
)

func (c Currency) String() string {
	switch c {
	case JPY:
		return "JPY"
	case USD:
		return "USD"
	default:
		return ""
	}
}

type Money struct {
	amount   decimal.Decimal
	currency Currency
}

func NewMoney(amount decimal.Decimal, currency Currency) *Money {
	return &Money{amount: amount, currency: currency}
}

func (m *Money) String() string {
	return fmt.Sprintf("%s%s", m.amount.String(), m.currency.String())
}

func (m *Money) Add(other *Money) (*Money, error) {
	if m.currency != other.currency {
		return nil, errors.New("invalid currency")
	}

	return NewMoney(m.amount.Add(other.amount), m.currency), nil
}

func (m *Money) Equals(other *Money) bool {
	// if m.currency != other.currency {
	// 	return false
	// }

	// return m.amount.Equal(other.amount)
	return reflect.DeepEqual(m, other)
}
