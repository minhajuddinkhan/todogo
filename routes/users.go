package routes

import (
	"fmt"
	"net/http"

	"github.com/darahayes/go-boom"

	conf "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/store"
	"github.com/minhajuddinkhan/todogo/utils"
)

//GetUserCSV GetUserCSV
func GetUserCSV(conf *conf.Configuration, store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		dir := "./csvs"
		filename := "/users.csv"

		users := []models.User{}
		store.GetAllUsers(&users)
		records := [][]string{}
		for _, user := range users {
			records = append(records, []string{
				user.Name,
				user.Address,
			})
		}

		err := utils.WriteCsv(dir, filename, records)
		if err != nil {
			boom.Forbidden(w, err.Error())
		}
		utils.Respond(w, fmt.Sprintf("filepath: %s", dir+filename))
	}
}

//RegisterUserRoutes RegisterUserRoutes
func RegisterUserRoutes(R router.RouterConf, conf *conf.Configuration, store store.Store) {
	R.RegisterHandlerFunc("POST", "/users/csv", GetUserCSV(conf, store))
}
