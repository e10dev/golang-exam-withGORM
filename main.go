package main

import (
	"e10dev.example/service/utils"
	"e10dev.example/service/webserver"

	"github.com/gin-gonic/gin"
)

func initMain() *gin.Engine {
	utils.DbInitTable()
	r := webserver.InitRouter()

	return r
}

func main() {
	r := initMain()

	r.Run()
}
