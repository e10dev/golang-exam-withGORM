package account

import (
	"e10dev.example/service/config"
	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func GetAccount(c *gin.Context) {
	param := c.Param("seq")
	var user model.User

	db := utils.DbConnect()

	res := db.Where("seq = ?", param).Take(&user)
	if res.Error != nil {
		if res.RowsAffected == 0 {
			utils.Response(c, config.BAD_REQUEST, "No User Found.")
			return
		}
		utils.Response(c, config.INTERNAL_SERVER_ERROR, "Query Error.")
		return
	}

	utils.ResponseJSON(c, user)
}
