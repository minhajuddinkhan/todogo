package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/darahayes/go-boom"
	"github.com/jinzhu/gorm"

	"github.com/gorilla/mux"

	conf "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/router"

	"github.com/minhajuddinkhan/todogo/models"
	"github.com/minhajuddinkhan/todogo/store"
	"github.com/minhajuddinkhan/todogo/utils"
)

//TodoRoutes TodoRoutes
func TodoRoutes(c *conf.Configuration, store store.Store) []router.Route {

	return []router.Route{
		{
			Method: "GET", URI: "/todos", Handler: GetTodos(c, store),
		},
		{
			Method: "GET", URI: "/todos/{id}", Handler: GetTodoByID(c, store),
		},
		{
			Method: "POST", URI: "/todos", Handler: CreateTodo(c, store),
		},
	}
}

//GetTodos GetTodos
func GetTodos(c *conf.Configuration, store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		var todos []models.Todo
		err := store.GetTodos(&todos).Error
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

}

//GetTodoByID GetTodoByID
func GetTodoByID(c *conf.Configuration, store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		todoID := vars["id"]
		var todo models.Todo

		err := store.GetTodoByID(&todo, todoID).Error
		if err != nil {
			if gorm.IsRecordNotFoundError(err) {
				boom.NotFound(w, "todo not found")
				return
			}
			boom.Internal(w, err)
			return
		}

		utils.Respond(w, todo)

	}
}

func CreateTodo(c *conf.Configuration, store store.Store) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		b, err := ioutil.ReadAll(r.Body)
		if err != nil {
			boom.Internal(w, err)
			return
		}

		var todo models.Todo
		if err := json.Unmarshal(b, &todo); err != nil {
			boom.BadRequest(w, err.Error())
			return
		}

		if err := store.CreateTodo(&todo).Error; err != nil {
			boom.BadRequest(w, err)
			return
		}
		utils.Respond(w, todo)

	}
}
