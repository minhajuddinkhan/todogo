package middlewares

import (
	"context"
	"net/http"

	"github.com/darahayes/go-boom"
	jwt "github.com/minhajuddinkhan/todogo/jsonwebtoken"
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

		decoded, err := jwt.DecodeJWT(Header, SecretKey)
		if err != nil {
			boom.BadRequest(w, err.Error())
			return
		}

		ctx := context.WithValue(r.Context(), Header, decoded)
		next.ServeHTTP(w, r.WithContext(ctx))

	}

}
