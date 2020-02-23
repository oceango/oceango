package main

import (
	"context"
	"github.com/oceango/middleware"
	"github.com/oceango/router"
	middlewares "github.com/oceango/skeleton/middleware"
	"net/http"
)

type OceanRouter interface {
	//AddRoute(httpMethod string, path string, handler interface{})
	//AddGroup(prefix string, func())
	Get(route string, handlerFunc http.HandlerFunc)
	Post(route string, handlerFunc http.HandlerFunc)
}

type OceanRoute struct {
	router *router.Router
}

func newOceanRoute() *OceanRoute {
	return &OceanRoute{router:router.New()}
}

func (o OceanRoute) Get(path string, handlerFunc http.HandlerFunc) {
	o.handleAdaptor(http.MethodGet, path, handlerFunc)
}

func (o OceanRoute) Post(path string, handlerFunc http.HandlerFunc) {
	o.handleAdaptor(http.MethodPost, path, handlerFunc)
}


func (o OceanRoute) handleAdaptor(httpMethod string, path string, handleFunc http.HandlerFunc ) {
	var handle router.Handle
	commonHandlers := middleware.Use(middlewares.AuthMiddleware, middlewares.LoggingMiddleware, middlewares.RecoverMiddleware)
	handle = wrapHandler(commonHandlers.ThenFunc(handleFunc))

	o.router.Handle(httpMethod, path, handle)
}

func wrapHandler(h http.Handler) router.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps router.Params) {
		ctx := r.Context()
		ctx = context.WithValue(ctx, "params", ps)
		h.ServeHTTP(w, r.WithContext(ctx))
	}
}




