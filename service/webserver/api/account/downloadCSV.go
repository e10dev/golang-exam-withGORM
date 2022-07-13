package account

import (
	"bufio"
	"bytes"
	"encoding/csv"
	"strconv"

	"e10dev.example/service/model"
	"e10dev.example/service/utils"
	"github.com/gin-gonic/gin"
)

func DownloadCSV(c *gin.Context) {
	var users []model.User

	db := utils.DbConnect()
	if res := db.Find(&users); res.Error != nil {
		utils.Response(c, utils.INTERNAL_SERVER_ERROR, "Query Error.")
		panic(res.Error)
	}

	b := &bytes.Buffer{}
	w := csv.NewWriter(bufio.NewWriter(b))

	header := []string{"seq", "id", "pw", "name", "email", "hp", "role", "state", "description"}

	if err := w.Write(header); err != nil {
		utils.Response(c, utils.INTERNAL_SERVER_ERROR, "Write CSV Error.")
		panic(err)
	}

	for _, u := range users {
		println(u.Seq, u.Id, u.Pw, u.Name, u.Email, u.Hp, u.Role, u.State, u.Description)
		var record []string

		record = append(record, strconv.Itoa(u.Seq))
		record = append(record, u.Id)
		record = append(record, u.Pw)
		record = append(record, u.Name)
		record = append(record, u.Email)
		record = append(record, u.Hp)
		record = append(record, strconv.Itoa(u.Role))
		record = append(record, strconv.Itoa(u.State))
		record = append(record, u.Description)
		if err := w.Write(record); err != nil {
			utils.Response(c, utils.INTERNAL_SERVER_ERROR, "Write CSV Error.")
			panic(err)
		}
	}
	w.Flush()

	utils.ResponseCSVFile(c, b.Bytes())
}
