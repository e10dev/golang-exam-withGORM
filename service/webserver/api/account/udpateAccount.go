package account

import (
	"e10dev.example/service/config"
	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	var user model.User

	//seq := c.Param("seq")
	if err := c.BindJSON(&user); err != nil {
		utils.Response(c, config.BAD_REQUEST, "Bad Params.")
		panic(err)
	}

	//db := utils.DbConnect()

	//db.Model(&user{})
}
