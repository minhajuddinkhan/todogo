package routes

import (
	conf "github.com/minhajuddinkhan/todogo/config"
	"github.com/minhajuddinkhan/todogo/router"
	"github.com/minhajuddinkhan/todogo/store"
)

//RegisterAllRoutes RegisterAllRoutes
func RegisterAllRoutes(r router.RouterConf, c *conf.Configuration, store store.Store) {

	r.RegisterMultipleHandlers(TodoRoutes(c, store))
	r.RegisterMultipleHandlers(AuthRoutes(c, store))
	r.RegisterMultipleHandlers(UserRoutes(c, store))
}
