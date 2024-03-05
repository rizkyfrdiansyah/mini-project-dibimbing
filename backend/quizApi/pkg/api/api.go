package api

import (
	"log"

	"github.com/gin-gonic/gin"
)

type (
	Api struct {
		server *gin.Engine
		routers []Router
	}

	Router interface {
		Route(handler *gin.RouterGroup)
	}
)

func (a Api) Start() error {
	root := a.server.Group("/")
	for _, router := range a.routers {
		router.Route(root)
	}

	err := a.server.Run(":8080")
	if err != nil {
		log.Println(err)
		return err
	}
	return err
}