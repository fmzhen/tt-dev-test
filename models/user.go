package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	TYPE_USER = "user"
)

type User struct {
	Id   int    `orm:"column(id);pk;auto" json:"id"`
	Name string `orm:"column(name)" json:"name" valid:"Required"`
	Type string `orm:"-" json:"type"`
}

func (t *User) TableName() string {
	return "user"
}

func init() {
	orm.RegisterModel(new(User))
}

// AddUser insert a new User into database and returns
// last inserted Id on success.
func AddUser(m *User) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	m.Id = int(id)
	m.Type = TYPE_USER
	return
}

// GetAllUser retrieves all User. Returns empty list if no records exist
func GetAllUser() (allUser []User, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User))

	_, err = qs.All(&allUser)
	if err != nil {
		beego.Warning("get user failed ,err=", err)
		return allUser, err
	}
	for i := range allUser {
		allUser[i].Type = TYPE_USER
	}
	return allUser, err
}
