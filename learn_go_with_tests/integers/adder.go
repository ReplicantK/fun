package integers

import "strings"

func Add(num1, num2 int) int {

	return num1 + num2
}

func Repeat(text string, n_count int) string {
	var output strings.Builder

	for i := 0; i < n_count; i++ {
		output.WriteString(text)
	}

	return output.String()
}
