package reverse

func Reverse(input string) string {
	inputArray := []rune(input)
	inputLength := len(inputArray)
	for i := 0; i < inputLength/2; i++ {
		aux := inputArray[i]
		inputArray[i] = inputArray[inputLength-i-1]
		inputArray[inputLength-i-1] = aux
	}
	return string(inputArray)
}
