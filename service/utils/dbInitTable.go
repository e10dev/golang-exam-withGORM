package utils

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"e10dev.example/service/model"
)

func DbInitTable() {
	fp, err := os.Open("service/dummy_data/account.json")
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	byteValue, _ := ioutil.ReadAll(fp)

	var users []model.User

	if err = json.Unmarshal(byteValue, &users); err != nil {
		panic(err)
	}

	db := DbConnect()

	if db.Migrator().HasTable(&users) {
		if err = db.Migrator().DropTable(&users); err != nil {
			panic(err)
		}
	}

	if err = db.AutoMigrate(&model.User{}); err != nil {
		panic(err)
	}

	res := db.Create(&users)
	if res.Error != nil {
		panic(res.Error)
	}
	model.Seq = res.RowsAffected
}
