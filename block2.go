package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func getMD5Hash(message string) string {
	hash := md5.Sum([]byte(message))
	return hex.EncodeToString(hash[:])
}

func main() {
	mypassword := "secret"
	fmt.Println("MD5 Hashed value: ", getMD5Hash(mypassword))

}