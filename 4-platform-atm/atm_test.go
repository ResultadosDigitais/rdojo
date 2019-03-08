package atm

import "testing"
import "reflect"

func TestWithDraw(t *testing.T) {

	tests := []struct {
		Description   string
		Value         int
		ExpectedMoney map[int]int
		ExpectedError bool
	}{
		{
			Description:   "Sacar 100 reais",
			Value:         100,
			ExpectedMoney: map[int]int{100: 1},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 50 reais",
			Value:         50,
			ExpectedMoney: map[int]int{50: 1},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 10 reais",
			Value:         10,
			ExpectedMoney: map[int]int{10: 1},
			ExpectedError: false,
		},

		{
			Description:   "Sacar 20 reais",
			Value:         20,
			ExpectedMoney: map[int]int{20: 1},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 200 reais",
			Value:         200,
			ExpectedMoney: map[int]int{100: 2},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 80 reais",
			Value:         80,
			ExpectedMoney: map[int]int{50: 1, 20: 1, 10: 1},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 60 reais",
			Value:         60,
			ExpectedMoney: map[int]int{50: 1, 10: 1},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 42 reais",
			Value:         42,
			ExpectedMoney: map[int]int{},
			ExpectedError: true,
		},
		{
			Description:   "Sacar 1100 reais",
			Value:         1100,
			ExpectedMoney: map[int]int{100: 10, 50: 2},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 4100 reais",
			Value:         4100,
			ExpectedMoney: map[int]int{},
			ExpectedError: true,
		},
		{
			Description:   "Sacar 4000 reais",
			Value:         4000,
			ExpectedMoney: map[int]int{100: 10, 50: 20, 20: 50, 10: 100},
			ExpectedError: false,
		},
		{
			Description:   "Sacar 0 reais deveria retornar error",
			Value:         0,
			ExpectedMoney: map[int]int{},
			ExpectedError: true,
		},
		{
			Description:   "Sacar um valor negativo deveria retornar erro",
			Value:         -1000,
			ExpectedMoney: map[int]int{},
			ExpectedError: true,
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			atm := NewATM()
			result, err := atm.Withdraw(test.Value)

			if !reflect.DeepEqual(result, test.ExpectedMoney) {
				t.Errorf("Expected %v got %v", test.ExpectedMoney, result)
			}

			if err != nil && !test.ExpectedError {
				t.Errorf("Expected no error, got %v", err)
			}

			if test.ExpectedError && err == nil {
				t.Errorf("Expected error, got nil")
			}
		})
	}
}

func TestWithDrawStateful(t *testing.T) {
	tests := []struct {
		Description   string
		Values        []int
		ExpectedError bool
		ExpectedMoney map[int]int
	}{
		{
			Description:   "Can make a withdraw after other withdraw",
			Values:        []int{100, 100},
			ExpectedError: false,
			ExpectedMoney: map[int]int{100: 1},
		},
		{
			Description:   "Can make a withdraw after other withdraw after the end of bill type",
			Values:        []int{1000, 100},
			ExpectedError: false,
			ExpectedMoney: map[int]int{50: 2},
		},
		{
			Description:   "Cannot make a withdraw when ATM has no money",
			Values:        []int{4000, 300},
			ExpectedError: true,
			ExpectedMoney: map[int]int{},
		},
	}

	for _, test := range tests {
		t.Run(test.Description, func(t *testing.T) {
			atm := NewATM()
			var result map[int]int
			var err error
			for _, v := range test.Values {
				result, err = atm.Withdraw(v)
			}

			if !reflect.DeepEqual(result, test.ExpectedMoney) {
				t.Errorf("Expected %v got %v", test.ExpectedMoney, result)
			}

			if err != nil && !test.ExpectedError {
				t.Errorf("Expected no error, got %v", err)
			}

			if test.ExpectedError && err == nil {
				t.Errorf("Expected error, got nil")
			}
		})
	}
}
