package greetings

import (
	"errors"
	"fmt"
	"log"
)

func Hello(name string) (string, error) {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	if name == "" {
		log.Println("Name cannot be empty")
		return "", errors.New("Name cannot be empty")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}

func main() {
	message, err := Hello("daman")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	fmt.Println(message)

}
