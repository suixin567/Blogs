//商品模型
package models

import (
	"time"

	"github.com/astaxie/beego/orm"     //引入beego的orm
	_ "github.com/go-sql-driver/mysql" //引入beego的mysql驱动
)

//模型
type Contactus struct {
	Id          int       `orm:"pk"`
	Phone       string    `orm:"size(15)"`
	Message     string    `orm:"size(255)"`
	Count       int       `orm:"size(10)"`
	Createdtime time.Time `orm:"auto_now_add;type(datetime)"`
}

//增删改查
func (v *Contactus) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(v)
	return err
}

func (h *Contactus) Insert() (n int64, err error) {
	o := orm.NewOrm()
	n, err = o.Insert(h)
	return n, err
}

//func (v *) Update() (err error) {
//	o := orm.NewOrm()
//	_, err = o.Update(&v)
//	return err
//}

//func (v *Contactus) Delete() (err error) {
//	o := orm.NewOrm()
//	_, err = o.Delete(&v)
//	return err
//}

//收藏房间
func (v *Contactus) Save() (err error) {
	v.Id = 0
	v.Count = 0
	h8, _ := time.ParseDuration("8h")
	v.Createdtime = time.Now().Add(h8)
	_, err = v.Insert()
	return err
}
