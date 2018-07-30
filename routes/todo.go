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
	"github.com/minhajuddinkhan/todogo/utils"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {

	pg := r.Context().Value(constants.DbKey).(*db.PostgresDB)
	conn := pg.EstablishConnection()
	defer conn.Close()
	var todos []models.Todo
	conn.Preload("User").Find(&todos)
	utils.Respond(w, todos)

}

func GetTodoById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	todoID := vars["id"]
	pg := r.Context().Value(constants.DbKey).(*db.PostgresDB)
	conn := pg.EstablishConnection()
	defer conn.Close()

	var todo models.Todo
	err := conn.First(&todo, todoID).Error
	if gorm.IsRecordNotFoundError(err) {
		boom.NotFound(w, "todo not found")
		return
	}

	utils.Respond(w, todo)
}

//RegisterTodoRoutes RegisterTodoRoutes
func RegisterTodoRoutes(r router.RouterConf) {

	r.RegisterHandlerFunc("GET", "/todos", GetTodos)
	r.RegisterHandlerFunc("GET", "/todos/{id}", GetTodoById)
}
