package controllers

import (
	"io/ioutil"
	"os"
	"path"
	"time"

	"github.com/astaxie/beego"
	"github.com/dreamCodeMan/easy-fs/app/common"
	"github.com/dreamCodeMan/easy-fs/app/models"
)

type ApiController struct {
	beego.Controller
}

type JSON struct {
	Code  int         `json:"code"`
	Msg   string      `json:"msg"`
	Count int64       `json:"count,omitempty"`
	Data  interface{} `json:"data,omitempty"`
}

//把需要返回的结构序列化成json 输出
func (c *ApiController) ajaxResult(code int, message string, data interface{}) {
	c.Data["json"] = JSON{
		code,
		message,
		0,
		data,
	}
	c.ServeJSON()
	c.StopRun()
}

func (c *ApiController) List() {
	filePath := c.GetString("dir", "/")
	filePath = path.Join(models.RootPath, filePath)
	beego.Info("获取列表:", filePath)
	list, err := common.GetFilesAndDirs(filePath)
	if err != nil {
		c.ajaxResult(0, "获取服务端文件列表失败", nil)
	}
	c.ajaxResult(200, "success", list)
}

func (c *ApiController) Del() {
	filename := c.GetString("file", "/")
	filename = path.Join(models.RootPath, filename)
	beego.Info("删除文库:", filename)
	if err := os.Remove(filename); err != nil {
		c.ajaxResult(0, "获取服务端文件列表失败", nil)
	}
	c.ajaxResult(200, "文件删除成功", nil)
}

func (c *ApiController) Upload() {
	if c.Ctx.Request.Method != "POST" {
		c.ajaxResult(0, "非法请求", nil)
	}
	var err error
	file, fileHead, err := c.GetFile("file")
	if err != nil {
		c.ajaxResult(0, err.Error(), nil)
	}
	defer file.Close()

	filename := fileHead.Filename

	filedir := path.Join(models.RootPath, time.Now().Format("2006/01/02"))
	objFileName := path.Join(filedir, filename)
	tempFileName := path.Join(filedir, filename+".tmp")
	if common.Exist(objFileName) {
		c.ajaxResult(200, "文件上传成功", nil)
	}

	common.CreateDir(filedir)
	data, err := ioutil.ReadAll(file)
	if err != nil {
		c.ajaxResult(0, err.Error(), nil)
	}

	if err = ioutil.WriteFile(tempFileName, data, 0777); err != nil {
		c.ajaxResult(0, err.Error(), nil)
	}
	os.Rename(tempFileName, objFileName)
	c.ajaxResult(200, "文件上传成功", nil)
}
