package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Jwt() {
	//MapClaims Map
	mySigningKey := []byte("celeste")
	//StandardClaims struct
	c := MyClaims{
		Username: "celst",
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: time.Now().Unix() + 60*60*2,
			Issuer:    "celst",
		},
	}
	//t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	//	"exp":      time.Now().Unix() + 5,
	//	"iss":      "celeste",
	//	"nbf":      time.Now().Unix() - 5,
	//	"username": "celeste",
	//})
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	s, err := t.SignedString(mySigningKey)
	if err != nil {
		fmt.Println("&s", err)
	} else {
		fmt.Println(s)
		//token, err := jwt.ParseWithClaims(s, &jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		//	return mySigningKey, nil
		//})
		token, err := jwt.ParseWithClaims(s, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
			return mySigningKey, nil
		})
		if err != nil {
			fmt.Println("err===>", err)
			return
		}
		//fmt.Println("token===>", token.Claims.(*jwt.MapClaims))
		fmt.Println(token.Claims.(*MyClaims))
	}
}
