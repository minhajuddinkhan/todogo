package context

import (
	"context"
	"fmt"
	"net/http"
)

type Context struct {
	context.Context
}

func NewContext(ctx context.Context) Context {
	return Context{ctx}
}
func GetContext(r *http.Request) Context {

	ctx, ok := r.Context().(Context)
	if !ok {
		fmt.Println("Failed to get custom Context from request")
		return ctx
	}
	return ctx
}

//With With
func (ctx *Context) With(payload interface{}, key int) Context {

	return Context{context.WithValue(ctx, key, payload)}
}

//WithFunc WithFunc
func (ctx *Context) WithFunc(fn interface{}, key int) Context {

	return Context{context.WithValue(ctx, key, fn)}
}
