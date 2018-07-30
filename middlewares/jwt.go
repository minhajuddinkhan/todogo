package middlewares

import (
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/dgrijalva/jwt-go"
	"github.com/minhajuddinkhan/todogo/context"
)

//AuthenticateJWT AuthenticateJWT
func AuthenticateJWT(keyForDecodedDataAccess int, SecretKey string, Header string) func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

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
			boom.Internal(w, err)
			return
		}

		ctx := context.GetContext(r)
		ctx.With(decoded, keyForDecodedDataAccess)
		r.WithContext(ctx)

	}

}
