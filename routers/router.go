// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"tt-dev-test/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/users",
		&controllers.UserController{}, "GET:GetAll")
	beego.Router("/users",
		&controllers.UserController{}, "POST:AddUser")

	beego.Router("/users/:user_id/relationships",
		&controllers.RelationshipController{}, "GET:GetRelationship")
	beego.Router("/users/:user_id/relationships/:other_user_id",
		&controllers.RelationshipController{}, "PUT:UpdateRelationship")
}
