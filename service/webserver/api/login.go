package api

import (
	"e10dev.example/service/config"
	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var user model.UserCompact
	if err := c.BindJSON(&user); err != nil {
		utils.Response(c, config.BAD_REQUEST, "User Info Bad Request.")
		panic(err)
	}

	var result model.UserCompact
	// 이렇게도 사용 가능
	// test := map[string]interface{}{}

	db := utils.DbConnect()
	res := db.Table("users").Where("id = ?", user.Id).Take(&result)
	if res.Error != nil {
		utils.Response(c, config.INTERNAL_SERVER_ERROR, "Query Error.")
		return
	}

	if result.Id != "" {
		if result.Pw == user.Pw {
			utils.Response(c, config.OK, "Ok.")
			return
		}
		utils.Response(c, config.BAD_REQUEST, "Password is Wrong.")
		return
	}
	utils.Response(c, config.BAD_REQUEST, "Cannot find User.")
}
