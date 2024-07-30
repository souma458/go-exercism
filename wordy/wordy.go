package wordy

import (
	"strconv"
	"strings"
	"unicode"
)

var operations = []string{"plus", "minus", "multiplied by", "divided by"}
var additionOpLength = len(operations[0])
var subtractionOpLength = len(operations[1])
var multiplicationOpLength = len(operations[2])
var divisionOpLength = len(operations[3])
var questionStart = "What is"

func Answer(question string) (int, bool) {
	questionAux := strings.Replace(question, questionStart, "", 1)
	if question == questionAux || len(questionAux) < 3 {
		return 0, false
	}

	firstNumber, numberLength := nextNumber(0, questionAux)
	if numberLength == 0 {
		return 0, false
	}

	// no operation
	if questionAux[numberLength+1] == '?' {
		return firstNumber, true
	}

	return parseQuestion(questionAux[numberLength+1:], firstNumber)
}

func parseQuestion(question string, current int) (int, bool) {
	question = strings.TrimSpace(question)
	if question[0] == '?' {
		return current, true
	}

	switch {
	case isAdditionNext(question):
		nextNumber, length := nextNumber(additionOpLength, question)
		if length > 0 {
			return parseQuestion(question[additionOpLength+length+1:], current+nextNumber)
		}
	case isSubtractionNext(question):
		nextNumber, length := nextNumber(subtractionOpLength, question)
		if length > 0 {
			return parseQuestion(question[subtractionOpLength+length+1:], current-nextNumber)
		}
	case isMultiplicationNext(question):
		nextNumber, length := nextNumber(multiplicationOpLength, question)
		if length > 0 {
			return parseQuestion(question[multiplicationOpLength+length+1:], current*nextNumber)
		}
	case isDivisionNext(question):
		nextNumber, length := nextNumber(divisionOpLength, question)
		if length > 0 {
			return parseQuestion(question[divisionOpLength+length+1:], current/nextNumber)
		}
	}

	return 0, false
}

func nextNumber(operatorLength int, question string) (int, int) {
	questionWithoutOperator := question[operatorLength+1:]
	numberString := ""

	for i := 0; i < len(questionWithoutOperator); i++ {
		if unicode.IsDigit(rune(questionWithoutOperator[i])) || questionWithoutOperator[i] == '-' {
			numberString += string(questionWithoutOperator[i])
		} else {
			break
		}
	}

	number, _ := strconv.Atoi(numberString)
	return number, len(numberString)
}

func isAdditionNext(question string) bool {
	return len(question) > additionOpLength+1 && question[:additionOpLength] == operations[0]
}

func isSubtractionNext(question string) bool {
	return len(question) > subtractionOpLength+1 && question[:subtractionOpLength] == operations[1]
}

func isMultiplicationNext(question string) bool {
	return len(question) > multiplicationOpLength+1 && question[:multiplicationOpLength] == operations[2]
}

func isDivisionNext(question string) bool {
	return len(question) > divisionOpLength+1 && question[:divisionOpLength] == operations[3]
}
