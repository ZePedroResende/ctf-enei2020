package main

import (
	"bufio"
	"fmt"
	"os"
	"unsafe"
)

func main() {

	fmt.Print("Enter Password: ")
	user_input, _, err := bufio.NewReader(os.Stdin).ReadLine()

	if err != nil {
		fmt.Println("Something is wrong with your computer, ", err)
	}

	const (
		EAX = uint8(unsafe.Sizeof(true))
	)

	var str []byte

	str = append(str, ((EAX<<EAX<<EAX^EAX)<<EAX<<EAX|EAX)<<EAX<<EAX)
	str = append(str, ((EAX<<EAX^EAX)<<EAX<<EAX|EAX)<<EAX<<EAX<<EAX)
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX ^ EAX))
	str = append(str, EAX<<EAX<<EAX<<EAX<<EAX<<EAX)
	str = append(str, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX<<EAX | EAX))
	str = append(str, ((((EAX<<EAX^EAX)<<EAX<<EAX|EAX)<<EAX^EAX)<<EAX^EAX)<<EAX)
	str = append(str, ((((EAX<<EAX^EAX)<<EAX|EAX)<<EAX<<EAX<<EAX^EAX)<<EAX ^ EAX))
	str = append(str, (((((EAX<<EAX^EAX)<<EAX|EAX)<<EAX<<EAX^EAX)<<EAX^EAX)<<EAX ^ EAX))
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX ^ EAX))
	str = append(str, (((EAX<<EAX^EAX)<<EAX|EAX)<<EAX<<EAX<<EAX^EAX)<<EAX)
	str = append(str, EAX<<EAX<<EAX<<EAX<<EAX<<EAX)
	str = append(str, (((EAX<<EAX^EAX)<<EAX|EAX)<<EAX<<EAX^EAX)<<EAX<<EAX)
	str = append(str, (((((EAX<<EAX^EAX)<<EAX<<EAX|EAX)<<EAX^EAX)<<EAX^EAX)<<EAX ^ EAX))
	str = append(str, EAX<<EAX<<EAX<<EAX<<EAX<<EAX)
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX|EAX)<<EAX^EAX)<<EAX<<EAX)
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX|EAX)<<EAX<<EAX<<EAX ^ EAX))
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX|EAX)<<EAX^EAX)<<EAX)
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX ^ EAX))
	str = append(str, EAX<<EAX<<EAX<<EAX<<EAX<<EAX)
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX|EAX)<<EAX<<EAX<<EAX ^ EAX))
	str = append(str, ((((EAX<<EAX^EAX)<<EAX|EAX)<<EAX<<EAX<<EAX^EAX)<<EAX ^ EAX))
	str = append(str, EAX<<EAX<<EAX<<EAX<<EAX<<EAX)
	str = append(str, ((EAX<<EAX^EAX)<<EAX<<EAX<<EAX<<EAX|EAX)<<EAX)
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX ^ EAX))
	str = append(str, (((EAX<<EAX^EAX)<<EAX<<EAX<<EAX|EAX)<<EAX<<EAX ^ EAX))
	str = append(str, (((EAX<<EAX^EAX)<<EAX|EAX)<<EAX<<EAX<<EAX^EAX)<<EAX)

	if string(user_input) == string(str) {
		fmt.Println(a)
	} else {
		fmt.Println("That's all folks !")
	}
}
