package main

import (
	"github.com/oceango/skeleton/controller"
	"net/http"
)

type routes []route

type route struct {
	path string
	handler http.Handler
}

func GetRoutes(router *OceanRoute) *OceanRoute {
	router.Post("/auth/login", controller.NewUserController().Login)
	router.Post("/auth/register", controller.NewUserController().Register)
	router.Get("/auth/info",controller.NewUserController().Info)
	router.Get("/welcome", controller.NewUserController().Welcome)
	return router
}
