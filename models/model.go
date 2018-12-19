package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id        int
	Name      string `orm:"size(50)"`
	UserName  string `orm:"size(50)"`
	PassWord  string `orm:"size(255)"`
	LastLogin time.Time
}

func init() {
	orm.RegisterModel(new(User))
	orm.RunSyncdb("default", false, true)

}

func UserLogin(username, password string) (user *User, err error) {
	o := orm.NewOrm()
	pwd := String2md5(password)
	user = &User{UserName: username, PassWord: pwd}
	err = o.Read(user, "UserName", "PassWord")
	fmt.Println(err)
	user.LastLogin = time.Now()
	o.Update(user)
	return user, err
}

func CreateUser(username, password, name string) (user *User, err error) {
	pwd := String2md5(password)
	user = &User{UserName: username, PassWord: pwd, Name: name, LastLogin: time.Now()}
	o := orm.NewOrm()
	_, err = o.Insert(user)
	fmt.Println(err)
	return user, err
}

func String2md5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has) //将[]byte转成16进制
}
