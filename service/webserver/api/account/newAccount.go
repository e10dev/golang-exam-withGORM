package account

import (
	"e10dev.example/service/config"
	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func NewAccount(c *gin.Context) {
	var user model.User

	if err := c.BindJSON(&user); err != nil {
		utils.Response(c, config.BAD_REQUEST, "Bad Params.")
		panic(err)
	}

	db := utils.DbConnect()

	var check model.User

	if db.Where("id = ?", user.Id).Take(&check).RowsAffected != 0 {
		utils.Response(c, config.BAD_REQUEST, "User Already Exists.")
		return
	}

	model.Seq++
	user.Seq = model.Seq
	if res := db.Create(&user); res.Error != nil {
		utils.Response(c, config.INTERNAL_SERVER_ERROR, "Query Error.")
		panic(res.Error)
	}

	utils.Response(c, config.CREATED, "OK")
}
