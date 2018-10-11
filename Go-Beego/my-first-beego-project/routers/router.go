package routers

import (
	"my-first-beego-project/controllers"
	"my-first-beego-project/filters"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/employees", &controllers.FirstController{}, "get:GetEmployees")
	beego.Router("/dashboard", &controllers.FirstController{}, "get:Dashboard")
	beego.Router("/home", &controllers.SessionController{}, "get:Home")
	beego.Router("/login", &controllers.SessionController{}, "get:Login")
	beego.Router("/logout", &controllers.SessionController{}, "get:Logout")
	beego.Router("/employee", &controllers.FirstController{}, "get:GetEmployee")
	beego.Router("/getFromCache", &controllers.CacheController{}, "get:GetFromCache")
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)
}
