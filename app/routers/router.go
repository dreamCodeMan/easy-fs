package routers

import (
	"github.com/dreamCodeMan/easy-fs/app/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.AutoRouter(&controllers.MainController{})
	beego.AutoRouter(&controllers.ApiController{})
}
