package middlewares

import (
	"net/http"

	context "github.com/minhajuddinkhan/todogo/context"
)

func AppendDatabaseContext(key int, db interface{}) func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		ctx := context.GetContext(r)
		r.WithContext(ctx.With(db, key))

		next.ServeHTTP(w, r)

	}
}
