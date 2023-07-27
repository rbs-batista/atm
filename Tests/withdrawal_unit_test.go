package services

import (
	services "atm/Services"
	"reflect"
	"testing"
)

func TestCalculateBills(t *testing.T) {

	banknotes := []int{200, 100, 50, 20, 10, 5, 2}

	tests := []struct {
		amount   int
		expected map[int]int
	}{
		{amount: 21, expected: map[int]int{2: 3, 5: 1, 10: 1}},
		{amount: 33, expected: map[int]int{2: 4, 5: 1, 20: 1}},
		{amount: 9, expected: map[int]int{2: 2, 5: 1}},
		{amount: 13, expected: map[int]int{2: 4, 5: 1}},
		{amount: 27, expected: map[int]int{2: 1, 5: 1, 20: 1}},
	}

	for _, test := range tests {
		result, _ := services.CalculateBills(test.amount, banknotes)

		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("calculateBills returned %v, expected %v", result, test.expected)
		}
	}
}
