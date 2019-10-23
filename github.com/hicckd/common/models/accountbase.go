package models

type AccountInfo struct {
	AccountId string `bson:"accountid"` //id号

	Account  string `bson:"account"`  //登录名
	Password string `bson:"password"` //密码 如果是微信登录的  则没有密码
	Salt     string `bson:"salt"`     //主要用户md5加密 算出password

	Phone string `bson:"phone"` //电话 如果是电话登录的  则phone和account一样
	Email string `bson:"email"` //邮箱	如果是邮箱登录 则邮箱和account一样
	QQ    string `bson:"qq"`    //qq  如果是qq登录 则qq和 account一样

	LastLogin   int64    `bson:"lastlogin"`
	LastIp      string   `bson:"lastip"`
	CreateTime  int64    `bson:"createtime"`
	PromiseList []string `bson:"promises"`

	//数据信息
	RealName  string `bson:"realname"`  //真实姓名
	AvatarUrl string `bson:"avatarurl"` //头像
	City      string `bson:"city"`      //城市
	Province  string `bson:"province"`  //省份
	Country   string `bson:"country"`   //国家
	Gender    int    `bson:"gender"`    //性别 1男 0女

}
