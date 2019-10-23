package models

import (
	"github.com/astaxie/beego/cache"
	"github.com/astaxie/beego"
)

var (
	CacheMgr cache.Cache
)

func init() {

	collectionName := beego.AppConfig.String("redisurl")
	connport := beego.AppConfig.String("redisport")
	dbNum := beego.AppConfig.String("redisdbNum")

	connectStr := `{`
	connectStr += `"key":"`+collectionName+`",`
	connectStr += `"conn":"`+connport+`",`
	connectStr += `"dbNum":"`+dbNum+`",`
	connectStr += `"password":`+""+`""}`

	cache, err := cache.NewCache("memory",connectStr )
	CacheMgr = cache
	if err != nil {
		beego.Error("连接redis错误 : ",err.Error())
		return
	}
	beego.Info("[redis.go] init : redis init scuess")
}