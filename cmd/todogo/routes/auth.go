package routes

import (
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"

	conf "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/store"

	"github.com/darahayes/go-boom"
	router "github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/utils"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Authorization struct {
	Authorization string
}

//UserLogin UserLogin
func UserLogin(conf *conf.Configuration, store *store.PgStore) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		requestPayload := &LoginRequest{}
		err := utils.DecodeRequest(r, requestPayload)
		if err != nil {
			boom.BadRequest(w, "Unable to parse json body.")
		}

		user := &models.User{
			Name:     requestPayload.Username,
			Password: requestPayload.Password,
		}
		err = store.GetUser(user).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				boom.NotFound(w, "user not found")
				return
			}
			boom.BadRequest(w, err.Error())
			return

		}
		signer := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"name": user.Name,
			"exp":  time.Now().Add(time.Minute * 20).Unix(),
		})
		token, err := signer.SignedString([]byte(conf.JWTSecret))
		if err != nil {
			boom.BadImplementation(w, "Could not sign JWT token")
		}

		auth := &Authorization{
			Authorization: token,
		}

		utils.Respond(w, auth)

	}

}

//RegisterAuthRoutes RegisterAuthRoutes
func RegisterAuthRoutes(R router.RouterConf, conf *conf.Configuration, store *store.PgStore) {
	R.RegisterHandlerFunc("POST", "/login", UserLogin(conf, store))
}
