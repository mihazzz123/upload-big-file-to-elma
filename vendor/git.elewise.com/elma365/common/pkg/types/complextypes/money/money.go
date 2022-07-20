package money

//go:generate ../../../../tooling/bin/easyjson money.go

import (
	"strconv"
	"strings"

	"git.elewise.com/elma365/common/pkg/types/complextypes/money/currency"
)

// Money is currency cents pair
//
// easyjson:json
type Money struct {
	C currency.Currency `json:"currency"`
	V int64             `json:"cents"`
}

// String выводит результат вычисления валюты, на базе разряда копеек
func (m Money) String() string {
	var result string

	str := strconv.FormatInt(m.V, 10) // строковое представление числа копеек

	if m.C.Units == 0 {
		result = str
	} else {
		diff := len(str) - int(m.C.Units)

		// если копеек меньше, чем разрядность валюты, добьем их нулями спереди
		if diff <= 0 {
			str = strings.Repeat("0", -diff+1) + str
		}

		index := len(str) - int(m.C.Units)

		result = str[:index] + "." + str[index:]
	}

	return result + " " + m.C.Code
}
