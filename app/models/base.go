package models

import (
	"github.com/astaxie/beego"
)

var (
	RootPath string
)

func init() {
	RootPath = beego.AppConfig.String("rootPath")
}
