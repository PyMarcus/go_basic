package main

// criar modulo go mod init [nome]
// instalar com go get "golang.org/x/crypto/bcrypt"
import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := "password"
	fmt.Println("PW BEFORE: ", password)
	// convert
	byteSliceHash, err := bcrypt.GenerateFromPassword([]byte(password), 10) // cost 4 to 31
	if err != nil {
		fmt.Println("ERRO: ", err)
	}
	fmt.Println("PW AFTER (Hash) ", string(byteSliceHash))

	// compare
	result := bcrypt.CompareHashAndPassword(byteSliceHash, []byte(password))
	if result == nil {
		fmt.Println("Correct!")
	} else {
		fmt.Println("Incorrect")
	}
}
