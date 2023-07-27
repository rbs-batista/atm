package services

import (
	"errors"
)

type Bill struct {
	Quantity int
	Banknote int
}

func CalculateBills(amount int, banknotes []int) (map[int]int, error) {

	bills := make(map[int]int)
	for _, banknote := range banknotes {
		if amount >= banknote {
			total := amount - banknote

			for total > 3 || total == 2 || total == 0 {
				bills[banknote] += 1
				amount -= banknote
				total -= banknote
			}
		}
	}

	if amount != 0 {
		return nil, errors.New("Valor inválido para saque!")
	}

	return bills, nil
}
