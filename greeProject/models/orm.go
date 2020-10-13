package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Id   int
	Name string
	Sex  int8
}

func init() {
	err:=orm.RegisterDataBase("default", "mysql", "root:root@tcp(127.0.0.1:3306)/gotest?charset=utf8", 30)
	if err != nil {
		fmt.Println("orm.RegisterDataBase() ERR!")
		return
	}
	orm.RegisterModel(new(User))
}

func PrintUserByORM() {
	o := orm.NewOrm()
	user := User{Id: 3}
	o.Read(&user)
	fmt.Println(user)
}
