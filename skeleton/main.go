package main

import (
	"github.com/oceango/web"
)


func main()  {
	web.BuildConfiguration()
	router := newOceanRoute()
	router = GetRoutes(router)
	application := web.NewApplication(router.router)
	application.Run()
}



