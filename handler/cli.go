package handler

import (
	"errors"
	"fmt"
	"log"

	"github.com/sahay-shashank/password-generator-golang/password"
)

func Cli(lengthptr *int) {
	length := *lengthptr
	if length < 7 {
		err := errors.New("'length' parameter needs to be greater than or equal to 7")
		log.Fatal(err)
		return
	}
	fmt.Printf("Generated Password: %s\n", password.Password(length))
}
