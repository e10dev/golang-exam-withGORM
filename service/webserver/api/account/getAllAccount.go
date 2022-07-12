package account

import (
	"e10dev.example/service/config"
	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func GetAllAccount(c *gin.Context) {
	var users []model.User

	db := utils.DbConnect()

	res := db.Find(&users)
	if res.Error != nil {
		utils.Response(c, config.INTERNAL_SERVER_ERROR, "Query Error.")
		return
	}

	utils.ResponseJSON(c, users)
}
