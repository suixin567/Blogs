package class

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type Article struct {
	Id      int    `orm:"pk"`
	Title   string `orm:"size(80)"`
	Content string `orm:"type(text)`
	Author  *User  `orm:"rel(fk)"`
	Replys  int
	Views   int

	Tags []*Tag `orm:"rel(m2m)"`

	Time time.Time `orm:"auto_now_add;type(datetime)"`

	Defunct bool
}

func (a *Article) ReadDB() (err error) {
	o := orm.NewOrm()
	if err = o.Read(a); err != nil {
		beego.Info(err)
	}
	_, _ = o.LoadRelated(a, "tags")
	return
}

func (a Article) Create() (n int64, err error) {
	o := orm.NewOrm()
	if n, err = o.Insert(&a); err != nil {
		beego.Info(err)
	}
	return
}

func (a Article) Update() (err error) {
	o := orm.NewOrm()
	if _, err = o.Update(&a); err != nil {
		beego.Info(err)
	}
	m2m := o.QueryM2M(&a, "Tags")
	//	m2m.Clear()
	//	m2m.Add(a.Tags)

	old := Article{Id: a.Id}
	_, _ = o.LoadRelated(&old, "Tags")

	//	insert
VI:
	for _, vi := range a.Tags {
		for _, vj := range old.Tags {
			if vi.Id == vj.Id {
				continue VI
			}
		}
		m2m.Add(vi)
	}

	//	del
VD:
	for _, vi := range old.Tags {
		for _, vj := range a.Tags {
			if vi.Id == vj.Id {
				continue VD
			}
		}
		m2m.Remove(vi)
	}

	return
}

func (a Article) Delete() (err error) {
	a.Defunct = true
	err = a.Update()
	return
}

//获取一个人的所有文章。（如果有标签，也会过滤标签）
func (a Article) Gets() (rets []Article) {
	o := orm.NewOrm()
	//	o.QueryTable("article").Filter("Author", a.Author).Filter("defunct", 0).All(&ret)
	//	return
	qs := o.QueryTable("article")

	if a.Author != nil {
		qs = qs.Filter("Author", a.Author)
	}

	if len(a.Tags) == 1 {
		qs = qs.Filter("Tags__Tag", a.Tags[0])
	}

	qs = qs.Filter("defunct", 0)

	//	Author
	qs = qs.RelatedSel()

	qs.All(&rets)

	//	Tags
	for i := range rets {
		_, _ = o.LoadRelated(&rets[i], "Tags")
	}
	return

}
