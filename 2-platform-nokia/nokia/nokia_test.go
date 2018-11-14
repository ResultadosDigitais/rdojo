package nokia

import (
	"reflect"
	"strings"
	"testing"
)

func TestGetTypedSequence(t *testing.T) {
	tests := []struct {
		Input          string
		ExpectedOutput string
		ExpectedError  error
	}{
		{
			Input:          "",
			ExpectedOutput: "",
		},

		{
			Input:          "A",
			ExpectedOutput: "2",
		},

		{
			Input:          "B",
			ExpectedOutput: "22",
		},

		{
			Input:          "C",
			ExpectedOutput: "222",
		},

		{
			Input:          "D",
			ExpectedOutput: "3",
		},

		{
			Input:          " ",
			ExpectedOutput: "0",
		},

		{
			Input:          "ABACATE",
			ExpectedOutput: "2_22_2_222_2833",
		},

		{
			Input:          "AA",
			ExpectedOutput: "2_2",
		},
		{
			Input:          "SEMPRE ACESSO O DOJOPUZZLES",
			ExpectedOutput: "77773367_7773302_222337777_777766606660366656667889999_9999555337777",
		},
		{
			Input:          "sempre ACESSO O DOJOPUZZLES",
			ExpectedOutput: "",
			ExpectedError:  UnsuportedCharacterError{},
		},
		{
			Input:          strings.Repeat("A", 260),
			ExpectedOutput: "",
			ExpectedError:  MaximumLengthError{},
		},
		{
			Input:          strings.Repeat("A", 255),
			ExpectedOutput: strings.Repeat("2_", 254) + "2",
		},
		{
			Input:          strings.Repeat(" ", 3),
			ExpectedOutput: "0_0_0",
		},
		{
			Input:          "SEBASTIÃO",
			ExpectedOutput: "",
			ExpectedError:  UnsuportedCharacterError{},
		},
		{
			Input:          "12349876",
			ExpectedOutput: "",
			ExpectedError:  UnsuportedCharacterError{},
		},
		{
			Input:          "ÕÉÃÔΩ$@",
			ExpectedOutput: "",
			ExpectedError:  UnsuportedCharacterError{},
		},
	}

	for _, test := range tests {
		t.Run("getTypedSequence "+test.Input, func(t *testing.T) {
			ret, err := getTypedSequence(test.Input)

			errorType := reflect.TypeOf(err)
			expectedErrorType := reflect.TypeOf(test.ExpectedError)

			if errorType != expectedErrorType {
				t.Fatalf("Error type must be %T but was %T", test.ExpectedError, err)
			}
			if test.ExpectedOutput != ret {
				t.Fatalf("Expected '%v' got '%v'", test.ExpectedOutput, ret)
			}
		})
	}
}
