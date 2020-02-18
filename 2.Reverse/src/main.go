package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	fmt.Print("Enter Password: ")
	user_input, _, err := bufio.NewReader(os.Stdin).ReadLine()

	if err != nil {
		fmt.Println("Something is wrong with your computer, ", err)
	}

	if string(user_input) == "The answer to life is beer" {
		fmt.Println(a)
	} else {
		fmt.Println("That's all folks !")
	}
}
