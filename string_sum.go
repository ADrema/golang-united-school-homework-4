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
	invalidExpression   = "invalid expression: %w"
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
		return "", fmt.Errorf(invalidExpression, errorEmptyInput)
	}
	sum := 0
	counter := 0
	value := ""
	symbol := ""
	var firstValue string
	var secondValue string

	runes := []rune(input)
	runesLength := len(runes)

	for i := 0; i < runesLength; i++ {
		runeValue := runes[i]
		stringValue := string(runeValue)
		isLast := i+1 == runesLength
		isOperator := runeValue == '+' || runeValue == '-'

		if counter == 2 {
			return "", fmt.Errorf(invalidExpression, errorNotTwoOperands)
		}

		if i == 0 {
			value = stringValue
			continue
		}

		if isLast {
			secondValue = symbol + value + stringValue
			continue
		}

		if isOperator && len(value) != 0 {
			firstValue = symbol + value
			value = ""
			symbol = stringValue
			counter++
			continue
		}
		value += stringValue
	}

	if counter == 0 {
		return "", fmt.Errorf(invalidExpression, errorNotTwoOperands)
	}

	valueOne, err := CheckValueISInteger(firstValue) // check if integer

	if err != nil {
		return "", fmt.Errorf("incorrect first operand: %w", err)
	}
	valueTwo, err := CheckValueISInteger(secondValue) // check if integer

	if err != nil {
		return "", fmt.Errorf("incorrect second operand: %w", err)
	}

	sum = valueOne + valueTwo
	output = strconv.Itoa(sum)
	return output, nil
}

func CheckValueISInteger(input string) (int, error) {
	value, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}
	return value, nil
}
