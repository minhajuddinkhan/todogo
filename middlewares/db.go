package middlewares

import (
	"net/http"

	"github.com/minhajuddinkhan/todogo/db"

	"github.com/minhajuddinkhan/todogo/context"
)

//AppendDatabaseContext AppendDatabaseContext
func AppendDatabaseContext(key int, database *db.PostgresDB) func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	return func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		ctx := context.NewContext(r.Context())
		ctx = ctx.With(database, key)
		next.ServeHTTP(w, r.WithContext(ctx))

	}
}
