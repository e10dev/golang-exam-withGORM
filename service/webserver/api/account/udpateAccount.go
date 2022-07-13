package account

import (
	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func UpdateAccount(c *gin.Context) {
	var body model.User

	seq := c.Param("seq")
	if err := c.BindJSON(&body); err != nil {
		utils.Response(c, utils.BAD_REQUEST, "Bad Params.")
		panic(err)
	}

	db := utils.DbConnect()

	if res := db.Model(&model.User{}).Where("seq = ?", seq).Updates(body); res.Error != nil {
		utils.Response(c, utils.INTERNAL_SERVER_ERROR, "Query Error.")
		panic(res.Error)
	}

	utils.Response(c, utils.CREATED, "OK")
}
