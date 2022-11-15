package main

import "fmt"

func threeFive(number int) string {
	switch {
	case number%5 == 0 && number%3 == 0:
		return "ThreeFive"
	case number%5 == 0:
		return "Five"
	case number%3 == 0:
		return "Three"
	default:
		// avoid returning error
		return fmt.Sprintf("%d", number)
	}
}
func main() {

	for i := 1; i <= 100; i++ {
		fmt.Println(threeFive(i))
	}
}
