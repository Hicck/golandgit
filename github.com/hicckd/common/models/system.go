package models

import "github.com/astaxie/beego"

type SystemConf struct {
	AesKey string
}

var (
	SystemConfigure SystemConf
)

func init() {
	SystemConfigure = SystemConf{}
	SystemConfigure.AesKey = beego.AppConfig.String("aeskey")
}

