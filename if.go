package main

import (
	"fmt"
	"strconv"
)

func main() {
	number1 := 17
	number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string

	if number1 > number2 {
		resultMessage = "number1 is bigger than number2 by " + strconv.Itoa(number1-number2)
		fmt.Println(resultMessage)
	} else {
		resultMessage = "number2 is bigger than number1 by " + strconv.Itoa(number2-number1)
		fmt.Println(resultMessage)
	}

}
