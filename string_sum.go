package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")

	if len(input) == 0 {
		return "", errorEmptyInput
	}
	sum := 0
	counter := 0
	value := ""
	symbol := "+"

	runes := []rune(input)
	runesLength := len(runes)

	for i := 0; i < runesLength; i++ {
		runeValue := runes[i]
		stringValue := string(runeValue)
		isLast := i+1 == runesLength
		isOperator := runeValue == '+' || runeValue == '-'

		if counter > 2 {
			return "", errorNotTwoOperands
		}

		if (isOperator && len(value) != 0) || isLast {
			if isLast {
				value = symbol + value + stringValue
			} else {
				value = symbol + value
			}
			numValue, err := CheckValueISInteger(value) // check if integer
			if err != nil {
				return "", err
			}
			sum += numValue
			value = ""
			symbol = stringValue
			counter++
			continue
		}
		if isOperator {
			symbol = stringValue
			continue
		}
		value += stringValue
	}

	return string(sum), nil
}

func CheckValueISInteger(input string) (int, error) {
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("incorrect value(s) are provided errors: %s", err)
	}
	return value, nil
}
