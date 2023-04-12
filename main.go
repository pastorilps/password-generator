package main

import (
	"fmt"

	"github.com/pastorilps/password-generator/password"
)

func main() {
	pass := password.GeneratePassword(8)
	cryptoPass := password.GetMd5Hash(pass)

	fmt.Println(pass)
	fmt.Println(cryptoPass)
}
