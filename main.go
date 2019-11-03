package main

import (
	_ "github.com/dreamCodeMan/easy-fs/app/routers"

	"github.com/astaxie/beego"
	"github.com/dreamCodeMan/easy-fs/app/models"
)

func main() {
	beego.SetStaticPath("/file", models.RootPath)
	beego.Run()
}
