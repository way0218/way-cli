package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	v1 "{{.Appname}}/api/product/app/v1"
	"{{.Appname}}/config/initializer"
	"{{.Appname}}/internal/server"
	"{{.Appname}}/internal/server/service"
)

var set = wire.NewSet(
	// domains
	server.Set,

	// common
	initializer.Set,
)

type services struct {
	article *service.Artcile
}

func (s *services) register(r gin.IRouter) {
	v1.RegisterBlogServiceHTTPServer(r, s.article)
}