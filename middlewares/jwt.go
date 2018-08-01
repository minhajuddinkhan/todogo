package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/dgrijalva/jwt-go"
)

//AuthenticateJWT AuthenticateJWT
func AuthenticateJWT(keyForDecodedDataAccess int, SecretKey string, Header string) func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		if r.URL.Path == "/login" {
			next.ServeHTTP(w, r)
			return
		}

		headers := r.Header.Get(Header)
		if len(headers) == 0 {
			boom.Unathorized(w, "No Authoriaztion Headers")
			return
		}

		decoded, err := jwt.Parse(headers, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
			return []byte(SecretKey), nil
		})

		if err != nil {
			boom.BadRequest(w, err.Error())
			return
		}

		fmt.Println("decoded", decoded)

		ctx := context.WithValue(r.Context(), Header, decoded)
		next.ServeHTTP(w, r.WithContext(ctx))

	}

}
