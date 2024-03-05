package api

import (
	"fmt"
	"log"
	"quizApi/pkg/db"

	"github.com/gin-gonic/gin"
)

func Default() *Api {
	server := gin.Default()
	

	_, err := db.Default()
	if err != nil {
		log.Println(err)
		panic(fmt.Sprintf("panic at db connection: %s", err.Error()))
	}
	// var routers = []Router{

	// }
	return &Api{
		server: server,
		routers: nil,
	}
}