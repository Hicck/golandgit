package models

import (
	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	RedisConn redis.Conn
}

var RedisMgr Redis

func init() {
	RedisMgr = Redis{}
	RedisMgr.Connect()
}

func (this *Redis) Connect() {
	var err error
	this.RedisConn, err = redis.Dial("tcp", beego.AppConfig.String("redisurl"))
	if err != nil {
		beego.Error("连接redis出现错误:", err.Error())
		panic(err)
	}
	beego.Info("init redis scuess...")
}
func (this *Redis) Set(key string, obj interface{}) {
	_, err := this.RedisConn.Do("SET", key, obj)
	if err != nil {
		beego.Error("redis保存数据出现错误,", key, err.Error())
	}
}

//秒为单位
func (this *Redis) SetWithTime(key string, obj interface{}, time int) {
	_, err := this.RedisConn.Do("SET", key, obj, "EX", time*60)
	if err != nil {
		beego.Error("redis保存数据出现错误,", key, err.Error(), "EX", time*60)
	}
}
func (this *Redis) GetString(key string) (reslut string, getErr error) {
	reslut, getErr = redis.String(this.RedisConn.Do("GET", key))
	if getErr != nil {
		beego.Error("redis未找到数据:", key)
		return
	}
	return
}
func (this *Redis) GetTTL(key string) (reslut int) {
	reply, err := this.RedisConn.Do("TTL", key)
	reslut, getErr := redis.Int(reply, err)
	if getErr != nil {
		beego.Error("redis未找到数据:", key)
		return
	}
	return
}
func (this *Redis) IsExist(key string) (is_key_exit bool) {
	is_key_exit, err := redis.Bool(this.RedisConn.Do("EXISTS", key))
	if err != nil {
		beego.Error(err.Error())
		return
	}
	return
}
