package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

)

func main() {
	mySigningKey := []byte("hzwy23")
	// Create the Claims
	claims := &jwt.StandardClaims{
		NotBefore: int64(time.Now().Unix() - 1000),
		ExpiresAt: int64(time.Now().Unix() + 1000),
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Println("签名后的token信息:", ss)
	t, err := jwt.Parse(ss, func(*jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		fmt.Println("parase with claims failed.", err)
		return
	}
	fmt.Println("还原后的token信息claims部分:", t.Claims)
}
