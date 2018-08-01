package routes

import (
	conf "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/store"
)

//RegisterAllRoutes RegisterAllRoutes
func RegisterAllRoutes(r router.RouterConf, c *conf.Configuration, store store.Store) {

	RegisterTodoRoutes(r, c, store)
	RegisterAuthRoutes(r, c, store)
}
