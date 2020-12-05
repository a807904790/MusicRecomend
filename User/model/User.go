package model

import (
	"MusicRecomend/proto"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang/protobuf/descriptor"
	"github.com/jinzhu/gorm"
)

type User struct {
	ID string
	Name string
	PassWord string
}

var Db *gorm.DB

func initDB()  {
	var err error
	Db, err = gorm.Open("mysql", proto.MysqlDsn)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	user := User{"123", "123", "123" }
	Db.Create(user)
}