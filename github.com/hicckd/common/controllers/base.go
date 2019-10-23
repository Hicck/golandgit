package controllers

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/astaxie/beego"
	"github.com/hicckd/common/models"
	"github.com/hicckd/common/utils"
	"io"
)

type BaseController struct {
	beego.Controller
}
type Request struct {
	Timestamp string      `json:"timestamp"`
	Id        string      `json:"id"`
	Token     string      `json:"token"`
	Uuid      string      `json:"uuid"`
	Data      interface{} `json:"data"`
}

func (this *BaseController) Check() (request Request, err error) {

	rdata := string(this.Ctx.Input.RequestBody)
	//如果传过来的加密字符串长度小于33的  md5(32位) + aes(至少一位)
	if len(rdata) <= 33 {
		err = errors.New("请勿进行网络攻击")
		this.StopRun()
		return request, err
	}
	md5 := rdata[0:32]
	encrypt := rdata[32:]

	//解密数据层
	decrypt, err := this.DecryptRequestBody([]byte(encrypt))
	if err != nil {
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "请勿攻击网络，开发不易，谢谢1",
			Data:    nil,
		})
		this.StopRun()
		return request, err
	}

	//转化位josn
	if err = json.Unmarshal([]byte(decrypt), &request); err != nil {
		this.SendResponse(models.Response{
			MsgCode: models.MSG_ErrParam,
			Msg:     "参数错误",
			Data:    "",
		})
		this.StopRun()
		return request, err
	}

	//检测md5
	allow := this.CheckTokenMD5(request, md5)
	if allow == false {
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "请勿攻击网络，开发不易，谢谢2",
			Data:    nil,
		})
		this.StopRun()
		return request, err
	}

	return request, nil
}
func (this *BaseController) CheckToken(request Request) (res bool) {
	//beego.Info(request)
	//获取根据id获取token
	reslut, err := models.RedisMgr.GetString("token:" + request.Id)

	if err != nil {
		res = false
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "登陆过期，请重新登陆1",
			Data:    nil,
		})
		this.StopRun()
		return
	}

	if request.Token != reslut {
		res = false
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "登陆过期，请重新登陆2",
			Data:    nil,
		})
		this.StopRun()
		return
	}
	return true

}
func (this *BaseController) CheckUuid(request Request) (res bool) {
	//获取根据id获取token
	reslut, err := models.RedisMgr.GetString("uuid:" + request.Id)
	if err != nil {
		res = false
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "登陆过期，请重新登陆",
			Data:    nil,
		})
		this.StopRun()
		return
	}

	if request.Uuid != reslut {
		res = false
		this.SendResponse(models.Response{
			MsgCode: models.MSG_NotAllow,
			Msg:     "登陆过期，请重新登陆",
			Data:    nil,
		})
		this.StopRun()
		return
	}
	return true
}
func (this *BaseController) CheckTokenAndUuid(request Request) (res bool) {
	//return this.CheckToken(request) && this.CheckUuid(request)
	//return this.CheckToken(request)
	return true
}
func (this *BaseController) DecryptRequestBody(encryptbytes []byte) (string, error) {
	encrypt := string(encryptbytes)
	decodeBytes, err := base64.StdEncoding.DecodeString(encrypt)
	if err != nil {
		return "", err
	}
	decrypt, err := utils.AesDecrypt(decodeBytes, []byte(models.SystemConfigure.AesKey))
	if err != nil {
		return "", err
	}
	return string(decrypt), nil
}
func (this *BaseController) SendResponse(response models.Response) {
	//先对resp里面的data进行加密
	resp, _ := json.Marshal(response)
	crypted := utils.AesEncrypt(string(resp), models.SystemConfigure.AesKey)
	cryptedTxt := base64.StdEncoding.EncodeToString(crypted)
	io.WriteString(this.Ctx.ResponseWriter, cryptedTxt)
	this.StopRun()
}
func (this *BaseController) SendErrorParam() {
	this.SendResponse(models.Response{
		MsgCode: models.MSG_ErrParam,
		Msg:     "参数错误",
		Data:    "",
	})
	this.StopRun()
}
func (this *BaseController) Encrpty(src string) string {
	crypted := utils.AesEncrypt(src, models.SystemConfigure.AesKey)
	cryptedTxt := base64.StdEncoding.EncodeToString(crypted)
	return cryptedTxt
}
func (this *BaseController) CheckTokenMD5(request Request, md5 string) bool {
	md52 := utils.Md5([]byte(request.Token + request.Timestamp))

	if md5 != md52 {
		return false
	}
	return true
}
