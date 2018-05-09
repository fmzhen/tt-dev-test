package controllers

import (
	"encoding/json"
	"errors"
	"tt-dev-test/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
)

// UserController operations for User
type UserController struct {
	beego.Controller
}

// Post ...
// @Title Post
// @Description create User
// @Param	body		body 	models.User	true		"body for User content"
// @Success 201 {int} models.User
// @Failure 400 param is invalid
// @router / [post]
func (c *UserController) AddUser() {
	var user models.User
	valid := validation.Validation{}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err == nil {
		if b, _ := valid.Valid(&user); !b {
			c.Ctx.Output.SetStatus(400)
			beego.Warning("param is invalid, name shouldn't empty")
			err = errors.New("param is invalid, name shouldn't empty")
			c.Data["json"] = err.Error()
			c.ServeJSON()
			return
		}

		if _, err := models.AddUser(&user); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = user
			beego.Info("add uder success, user=", user)
		} else {
			c.Ctx.Output.SetStatus(500)
			beego.Warning("add user failed. err=", err, " user=", user)
			c.Data["json"] = err.Error()
		}
	} else {
		c.Ctx.Output.SetStatus(400)
		beego.Warning("add user failed,param invalid. err=", err, " user=", user)
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get all User
// @Success 200 {object} models.User
// @Failure 403
// @router / [get]
func (c *UserController) GetAll() {
	l, err := models.GetAllUser()
	if err != nil {
		c.Ctx.Output.SetStatus(500)
		beego.Warning("get all users faild,err=", err)
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}
