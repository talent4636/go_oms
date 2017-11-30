package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"errors"
	"crypto/md5"
	"encoding/hex"
)

type User struct{
	Id int `pk:"auto"`
	Username string
	Password string
	Created  string //时间保存成 xxxx-xx-xx xx:xx:xx 的格式
	Mobile string `orm:"size(20);null"`
	Email string `orm:"size(40);null"`
	Neckname string `orm:"null"`
}

var timeLayout string = "2006-01-02 15:04:05"

func (u *User) AdminInit() bool{
	o:=orm.NewOrm()
	if cnt,_ := o.QueryTable("oms_user").Filter("username","admin").Count();cnt<1{
		//没有admin
		return false
	}else{
		return true
	}
}

func CreateAdmin(){
	u := new(User)
	u.Create("admin","admin123")
}

func (u *User) Create(username string, password string) (int64, error){
	o := orm.NewOrm()
	if len(username)<4 || len(password)<6 {
		return 0, errors.New("错误的用户名和密码长度")
	}
	objMd := md5.New()
	objMd.Write([]byte(password))
	pass := hex.EncodeToString(objMd.Sum(nil))
	var userNew User = User{
		Username:username,
		Password:pass,
		Created:time.Now().Format(timeLayout),
	}
	return o.Insert(&userNew)
}

func (u *User) Check(username string, password string) bool{
	objMd := md5.New()
	objMd.Write([]byte(password))
	pass := hex.EncodeToString(objMd.Sum(nil))
	o := orm.NewOrm()
	if cnt,_ := o.QueryTable("oms_user").Filter("username",username).
		Filter("password", pass).Count();cnt>0{
			return true
	}else{
		return false
	}
}