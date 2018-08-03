package routes

import (
	"fmt"
	"io/ioutil"
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

//UserFileUpload UserFileUpload
func UserFileUpload(conf *conf.Configuration, store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		file, handler, err := r.FormFile("uploadfile")
		if err != nil {
			boom.BadRequest(w, err.Error())
			return
		}

		if err != nil {
			boom.BadRequest(w, "not found whatever "+err.Error())
		}
		err = utils.FileUpload(file, handler.Filename, "./userfiles")
		if err != nil {
			boom.Conflict(w, err.Error())
		}

		b, err := ioutil.ReadFile("./userfiles/" + handler.Filename)
		if err != nil {
			boom.BadData(w, "cannot read the saved file:"+err.Error())
			return
		}
		w.Write(b)
		//		utils.Respond(w, "File upload done!")

	}
}

//RegisterUserRoutes RegisterUserRoutes
func RegisterUserRoutes(R router.RouterConf, conf *conf.Configuration, store store.Store) {
	R.RegisterHandlerFunc("POST", "/users/csv", GetUserCSV(conf, store))
	R.RegisterHandlerFunc("POST", "/users/upload", UserFileUpload(conf, store))
}
