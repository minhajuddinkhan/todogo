package routes

import (
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"

	"github.com/minhajuddinkhan/todogo/constants"
	"github.com/minhajuddinkhan/todogo/router"

	"github.com/minhajuddinkhan/todogo/db"

	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/store"
	"github.com/minhajuddinkhan/todogo/utils"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {

	pg := r.Context().Value(constants.DbKey).(*db.PostgresDB)
	var todos []models.Todo
	err := store.GetTodos(pg, &todos).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			boom.NotFound(w, "Todos Not Found")
			return
		}
		boom.Internal(w, "Something went wrong")
		return
	}
	utils.Respond(w, todos)

}

func GetTodoById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoID := vars["id"]
	var todo models.Todo

	pg := r.Context().Value(constants.DbKey).(*db.PostgresDB)
	err := store.GetTodoById(pg, &todo, todoID).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			boom.NotFound(w, "todo not found")
			return
		}
		boom.Internal(w, err)
	}

	utils.Respond(w, todo)
}

//RegisterTodoRoutes RegisterTodoRoutes
func RegisterTodoRoutes(r router.RouterConf) {

	r.RegisterHandlerFunc("GET", "/todos", GetTodos)
	r.RegisterHandlerFunc("GET", "/todos/{id}", GetTodoById)
}
