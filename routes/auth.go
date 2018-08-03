package routes

import (
	"net/http"

	"github.com/jinzhu/gorm"

	conf "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/store"

	"github.com/darahayes/go-boom"
	"github.com/minhajuddinkhan/todogo/jsonwebtoken"
	router "github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/utils"
)

//LoginRequest LoginRequest
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

//Authorization Authorization
type Authorization struct {
	Authorization string
}

//AuthRoutes AuthRoutes
func AuthRoutes(conf *conf.Configuration, store store.Store) []router.Route {

	return []router.Route{
		{
			Method: "POST", URI: "/login", Handler: UserLogin(conf, store),
		},
	}
}

//UserLogin UserLogin
func UserLogin(conf *conf.Configuration, store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		requestPayload := &LoginRequest{}
		err := utils.DecodeRequest(r, requestPayload)
		if err != nil {
			boom.BadRequest(w, "Unable to parse json body.")
		}

		user := models.User{
			Name:     requestPayload.Username,
			Password: requestPayload.Password,
		}

		err = store.GetUser(&user).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				boom.NotFound(w, "user not found")
				return
			}
			boom.BadRequest(w, err.Error())
			return

		}
		token, err := jsonwebtoken.EncodeJWT(user.Name, conf.JWTSecret)

		if err != nil {
			boom.BadImplementation(w, "Could not sign JWT token")
		}

		auth := &Authorization{
			Authorization: token,
		}

		utils.Respond(w, auth)

	}

}
