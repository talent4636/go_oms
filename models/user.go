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

//创建用户
func (u *User) New(user User) (int64, error){
	o := orm.NewOrm()
	objMd := md5.New()
	objMd.Write([]byte(user.Password))
	pass := hex.EncodeToString(objMd.Sum(nil))
	user.Password = pass
	if user.Created == ""{
		user.Created = time.Now().Format(timeLayout)
	}
	return o.Insert(&user)
}

func (u *User) Modify(id int, modifyData map[string]interface{}) (int64, error){
	o := orm.NewOrm()
	user := u.GetOne(map[string]interface{}{"id":id})
	for name,value := range modifyData{
		//user.name = value
		switch name {
		case "id":
			break
		case "password":
			if value==""{
				break
			}
			objMd := md5.New()
			objMd.Write([]byte(value.(string)))
			pass := hex.EncodeToString(objMd.Sum(nil))
			user.Password = pass
			break
		case "neckname":
			user.Neckname = value.(string)
			break
		case "mobile":
			user.Mobile = value.(string)
			break
		case "email":
			user.Email = value.(string)
			break
		default:
			break
		}
	}
	return o.Update(&user)
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

func (u *User) GetList(filters map[string]string, offset int, limit int) []User{
	o := orm.NewOrm()
	var data []User
	tmp := o.QueryTable("oms_user")
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	_,err := tmp.Offset(offset).Limit(limit).All(&data)
	if err!=nil{
		return nil
	}else{
		return data
	}
}

func (u *User) GetOne(filters map[string]interface{}) User{
	o := orm.NewOrm()
	var data User
	tmp := o.QueryTable("oms_user")
	for key,value := range filters{
		tmp = tmp.Filter(key,value)
	}
	_,err := tmp.All(&data)
	if err!=nil{
		return User{}
	}else {
		return data
	}
}