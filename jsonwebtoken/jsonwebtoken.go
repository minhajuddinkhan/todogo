package jsonwebtoken

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//EncodeJWT EncodeJWT
func EncodeJWT(name string, JWTSecret string) (string, error) {

	signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": name,
		"exp":  time.Now().Add(time.Minute * 20).Unix(),
	})

	s, _ := signer.SignedString([]byte(JWTSecret))
	fmt.Println(s)
	return signer.SignedString([]byte(JWTSecret))
}

//DecodeJWT DecodeJWT
func DecodeJWT(headers string, SecretKey string) (*jwt.Token, error) {

	return jwt.Parse(headers, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(SecretKey), nil
	})

}
