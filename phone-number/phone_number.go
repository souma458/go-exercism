package phonenumber

import (
	"regexp"
	"strings"
)

var numberRegex = regexp.MustCompile(`[^0-9]+`)

type InvalidNumberError struct{}

func (e *InvalidNumberError) Error() string {
	return "invalid number"
}

func Number(phoneNumber string) (string, error) {
	var aux = strings.ReplaceAll(numberRegex.ReplaceAllString(phoneNumber, ""), " ", "")

	if len(aux) < 10 || len(aux) > 11 {
		return "", &InvalidNumberError{}
	}

	if len(aux) == 11 {
		if aux[0] == '1' {
			aux = aux[1:]
		} else {
			return "", &InvalidNumberError{}
		}
	}

	if aux[0] < '2' || aux[3] < '2' {
		return "", &InvalidNumberError{}
	}

	return aux, nil
}

func AreaCode(phoneNumber string) (string, error) {
	number, nrError := Number(phoneNumber)
	if nrError != nil {
		return "", nrError
	}

	return number[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, nrError := Number(phoneNumber)
	if nrError != nil {
		return "", nrError
	}

	return "(" + number[0:3] + ") " + number[3:6] + "-" + number[6:], nil
}
