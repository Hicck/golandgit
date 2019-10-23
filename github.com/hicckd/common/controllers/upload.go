package controllers

import (
	"github.com/astaxie/beego"
	"github.com/hicckd/common/models"
	"github.com/hicckd/common/utils"
	"github.com/micro/go-micro/util/log"
	"os"
	"path"
	"strings"
)

type UploadController struct {
	BaseController
}

func (this *UploadController) UploadImgage() {
	//beego.Info("body:  ", this.Ctx.Request.Body)
	//beego.Info("上传图片")
	//image，这是一个key值，对应的是html中input type-‘file’的name属性值
	f, h, err := this.GetFile("image")
	if err != nil {
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "上传图片失败:" + err.Error(),
			Data:    nil,
		})
		beego.Info(err.Error())
		return
	}
	ext := path.Ext(h.Filename)
	//得到文件的名称
	fileName := utils.UniqueId() + ext
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}

	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//保存文件到指定的位置
	//static/uploadfile,这个是文件的地址，第一个static前面不要有/
	path := path.Join("static/uploadfile/images", fileName)
	this.SaveToFile("image", path)
	//显示在本页面，不做跳转操作
	this.SendResponse(models.Response{
		MsgCode: models.MSG_OK,
		Msg:     "上传图片成功",
		Data:    path,
	})
	this.StopRun()
}
func (this *UploadController) UploadApk() {
	//image，这是一个key值，对应的是html中input type-‘file’的name属性值
	f, h, _ := this.GetFile("apk")
	//得到文件的名称
	fileName := h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}

	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//保存文件到指定的位置
	//static/uploadfile,这个是文件的地址，第一个static前面不要有/
	path := path.Join("static/uploadfile/apk", fileName)
	err := this.SaveToFile("apk", path)
	//beego.Info(err)
	if err != nil {
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "上传apk失败:" + err.Error(),
			Data:    nil,
		})
	}
	//显示在本页面，不做跳转操作
	this.SendResponse(models.Response{
		MsgCode: models.MSG_OK,
		Msg:     "上传apk成功",
		Data:    path,
	})
	this.StopRun()
}
func (this *UploadController) UploadWgt() {
	//image，这是一个key值，对应的是html中input type-‘file’的name属性值
	f, h, _ := this.GetFile("wgt")
	//得到文件的名称
	fileName := h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}

	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//保存文件到指定的位置
	//static/uploadfile,这个是文件的地址，第一个static前面不要有/
	path := path.Join("static/uploadfile/wgt", fileName)
	err := this.SaveToFile("wgt", path)
	beego.Info(err)
	if err != nil {
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "上传wgt失败:" + err.Error(),
			Data:    nil,
		})
	}
	//显示在本页面，不做跳转操作
	this.SendResponse(models.Response{
		MsgCode: models.MSG_OK,
		Msg:     "上传wgt成功",
		Data:    path,
	})
	this.StopRun()
}
func (this *UploadController) UploadFile() {
	//log.Info("UploadFile")
	//image，这是一个key值，对应的是html中input type-‘file’的name属性值
	f, h, err := this.GetFile("file")
	if err != nil {
		log.Fatal("upload err : " + err.Error())
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "上传文件失败:" + err.Error(),
			Data:    nil,
		})
		return
	}

	//得到文件的名称
	fileName := h.Filename
	arr := strings.Split(fileName, ":")
	if len(arr) > 1 {
		index := len(arr) - 1
		fileName = arr[index]
	}
	//关闭上传的文件，不然的话会出现临时文件不能清除的情况
	f.Close()
	//.png
	ext := path.Ext(fileName)
	ext = strings.Replace(ext, ".", "", -1)
	dirPath := path.Join("static/uploadfile", ext)
	if utils.IsDir(dirPath) == false {
		os.MkdirAll(dirPath, os.ModePerm)
	}
	filepath := path.Join(dirPath, fileName)
	err = this.SaveToFile("file", filepath)
	if err != nil {
		log.Fatal("upload err : " + err.Error())
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "上传文件失败:" + err.Error(),
			Data:    nil,
		})
		return
	}
	this.SendResponse(models.Response{
		MsgCode: models.MSG_OK,
		Msg:     "上传文件成功",
		Data:    filepath,
	})
	this.StopRun()
}
