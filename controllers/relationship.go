package controllers

import (
	"encoding/json"
	"errors"
	"strconv"
	"tt-dev-test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// RelationshipController operations for Relationship
type RelationshipController struct {
	beego.Controller
}

func (c *RelationshipController) GetRelationship() {
	idStr := c.Ctx.Input.Param(":user_id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetRelationshipByUserId(id)
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = v
	}
	c.ServeJSON()
}

func (c *RelationshipController) UpdateRelationship() {
	userIdStr := c.Ctx.Input.Param(":user_id")
	otherUserIdStr := c.Ctx.Input.Param(":other_user_id")
	userId, _ := strconv.Atoi(userIdStr)
	otherUserId, _ := strconv.Atoi(otherUserIdStr)

	v := models.Relationship{}
	valid := validation.Validation{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		v.UserId = userId
		v.OtherUserId = otherUserId
		if b, _ := valid.Valid(&v); !b {
			c.Ctx.Output.SetStatus(400)
			beego.Warning("param is invalid,state should be liked or disliked")
			err = errors.New("param is invalid, state should be liked or disliked")
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}

		if newRel, err := models.UpdateRelationship(v); err == nil {
			beego.Info("add/update relationship success, newRel=", newRel)
			c.Data["json"] = newRel
		} else {
			c.Ctx.Output.SetStatus(500)
			beego.Warning("add/update relationship failed, err=", err, " relationShip=", v)
			c.Data["json"] = err.Error()
		}
	} else {
		beego.Warning("add/update relationship failed,param invalid. err=", err,
			" requestBody=", string(c.Ctx.Input.RequestBody))
		c.Ctx.Output.SetStatus(400)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
