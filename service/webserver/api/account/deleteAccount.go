package account

import (
	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func DeleteAccount(c *gin.Context) {
	seq := c.Param("seq")

	db := utils.DbConnect()

	if res := db.Delete(&model.User{}, seq); res.Error != nil {
		utils.Response(c, utils.INTERNAL_SERVER_ERROR, "Query Error.")
		panic(res.Error)
	}

	utils.Response(c, utils.NO_CONTENT, "Ok.")
}
