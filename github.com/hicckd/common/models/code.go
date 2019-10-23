package models

var (
	MSG_OK = 200		//正确处理
	MSG_NoPromise = 401 //权限不够
	MSG_Error = 500		//服务器内部错误
	MSG_NotAllow = 501	//不允许不明身份
	MSG_Repeat = 503 	//重复数据
	MSG_NotFind = 504 	//没有找到相关数据
	MSG_ErrParam = 505 	//参数错误
)

type Response struct {
	MsgCode int `json:"msgcode"`
	Msg string `json:"message"`
	Data interface{} `json:"data"`
}
