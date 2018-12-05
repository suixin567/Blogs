package models

import (
	"DDN_XS/models/class"
	"fmt"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {

	orm.RegisterDataBase("default", "mysql", "server2.0:furlogin123654@/ddnweb?charset=utf8", 30)
	orm.SetMaxIdleConns("default", 30) //设置数据库最大空闲连接
	orm.SetMaxOpenConns("default", 30) //设置数据库最大连接数
	orm.RegisterModel(new(class.User))
	orm.RunSyncdb("default", false, true)
	fmt.Println("注册数据库")
}
