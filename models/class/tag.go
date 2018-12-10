package class

import "github.com/astaxie/beego/orm"

import "DDN_XS/modules"

type Tag struct {
	Id       int64
	Name     string     `orm:"index"`
	Articles []*Article `orm:"reverse(many)"`
}

func (t Tag) Get() *Tag {
	o := orm.NewOrm()
	o.QueryTable("tag").Filter("Name", t.Name).One(&t)
	if t.Id == 0 { //正常情况不为0
		return nil
	}
	return &t
}

func (t Tag) GetOrNew() *Tag {
	o := orm.NewOrm()
	_, _, _ = o.ReadOrCreate(&t, "Name")
	return &t
}

var bscolor []string = []string{"success", "primary", "danger", "warning"}

func (t Tag) RandColor() string {
	return bscolor[modules.RandInt(4)]
}
