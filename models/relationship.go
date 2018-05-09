package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

const (
	TYPE_RELATIONSHIP = "relationship"
	LIKED             = "liked"
	DISLIKED          = "disliked"
	MATCHED           = "matched"
)

type Relationship struct {
	Id          int    `orm:"column(id);auto;pk" json:"-"`
	UserId      int    `orm:"column(user_id)" json:"-"`
	OtherUserId int    `orm:"column(other_user_id)" json:"user_id"`
	State       string `orm:"column(state)" json:"state" valid:"Match(liked|disliked)"`
	Type        string `orm:"-" json:"type"`
}

func (t *Relationship) TableName() string {
	return "relationship"
}

func init() {
	orm.RegisterModel(new(Relationship))
}

// GetRelationshipByUserId retrieves Relationship by User Id. Returns error if
// UserId doesn't exist
func GetRelationshipByUserId(userId int) (list []Relationship, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Relationship))

	_, err = qs.Filter("UserId", userId).All(&list)
	if err != nil {
		beego.Warning("find relationships by userid failed,userId=", userId, " err=", err)
		return list, err
	}
	for i := range list {
		list[i].Type = TYPE_RELATIONSHIP
	}
	return list, err
}

// UpdateRelationship updates Relationship
func UpdateRelationship(m Relationship) (newRelationship Relationship, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Relationship))
	state := m.State

	//
	otherRelation := Relationship{UserId: m.OtherUserId, OtherUserId: m.UserId}
	o.Read(&otherRelation, "UserId", "OtherUserId")
	if otherRelation.State == LIKED && m.State == LIKED {
		state = MATCHED
	}

	if errRead := o.Read(&m, "UserId", "OtherUserId"); errRead == orm.ErrNoRows {
		//新增
		m.State = state
		_, err = o.Insert(&m)

	} else {
		//更新
		m.State = state
		_, err = qs.Filter("UserId", m.UserId).Filter("OtherUserId", m.OtherUserId).
			Update(orm.Params{"State": m.State})
	}

	m.Type = TYPE_RELATIONSHIP
	if err != nil {
		beego.Warning("add/update relationship failed, err=", err, " relationShip=", m)
	}

	return m, err
}
