package models

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/orm"
	"strings"
	"time"
)

type User struct {
	Id        int
	Name      string `orm:"size(50)"`
	UserName  string `orm:"size(50)"`
	PassWord  string `orm:"size(255)"`
	LastLogin time.Time
}

type Article struct {
	Id         int
	Title      string `orm:"size(255)"`
	Content    string `orm:"size(3000)"`
	CreateTime time.Time
	User       *User      `orm:"rel(fk)"`       // 多对一(外键)
	Tag        []*Tag     `orm:"rel(m2m)"`      // 多对多
	Comment    []*Comment `orm:"reverse(many)"` // 设置一对多的反向关系
}

type Tag struct {
	Id         int
	Name       string `orm:"size(20)"`
	CreateTime time.Time
	Article    []*Article `orm:"reverse(many)"`
}

type Comment struct {
	Id         int
	Content    string `orm:"size(255)"`
	CreateTime time.Time
	User       *User    `orm:"rel(fk)"`
	Parent     *Comment `orm:"null;rel(fk)"` // 允许为null
	Article    *Article `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(User), new(Tag), new(Article), new(Comment))
	orm.RunSyncdb("default", false, true)
}

func UserOne(userId int) (*User, error) {
	user := User{Id: userId}
	o := orm.NewOrm()
	err := o.Read(&user)
	return &user, err
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

func CreateArticle(user *User, title string, content string, tagString string) error {
	tags_id := strings.Split(tagString, ",")
	o := orm.NewOrm()
	var tags []*Tag
	if len(tagString) > 0 {
		tag_count, err := o.QueryTable(new(Tag)).Filter("id__in", tags_id).All(&tags)
		if int(tag_count) != len(tags_id) {
			return err
		}
	}
	article := Article{User: user, Title: title, Content: content, Tag: tags, CreateTime: time.Now()}
	_, err := o.Insert(&article)
	return err
}

func CreateTag(name string) error {
	tag := Tag{Name: name, CreateTime: time.Time{}}
	o := orm.NewOrm()
	_, _, err := o.ReadOrCreate(&tag, "Name")
	return err
}

func ListArticle(pageNumber, pageSize int) ([]*Article, int64) {
	offset := pageNumber - 1*pageSize
	o := orm.NewOrm()
	var articles []*Article
	num, err := o.QueryTable(new(Article)).OrderBy("-create_time").Limit(pageSize, offset).All(&articles)
	fmt.Println(num)
	fmt.Println(err)
	return articles, num
}

func CreateComment(user *User, content string, articleId int, parentId int) error {
	article := Article{Id: articleId}
	comment := Comment{Content: content, User: user, Article: &article, CreateTime: time.Now()}
	if parentId != 0 {
		parent := Comment{Id: parentId}
		comment.Parent = &parent
	}
	o := orm.NewOrm()
	_, err := o.Insert(&comment)
	fmt.Println(err)
	return err
}

func ListComment(pageNumber, pageSize int) ([]*Comment, int64) {
	offset := pageNumber - 1*pageSize
	o := orm.NewOrm()
	var comments []*Comment
	num, err := o.QueryTable(new(Comment)).OrderBy("-create_time").Limit(pageSize, offset).All(&comments)
	fmt.Println(num)
	fmt.Println(err)
	return comments, num
}

func ListTag() []*Tag {
	o := orm.NewOrm()
	var tags []*Tag
	o.QueryTable(new(Comment)).OrderBy("-CreateTime").All(&tags)
	return tags
}

func Paginator(pageNumber, pageSize int) (int, int) {
	start := pageNumber - 1*pageSize
	end := pageNumber * pageSize
	return start, end
}

func ArticleOne(id int) *Article {
	article := new(Article)
	o := orm.NewOrm()
	o.QueryTable(article).Filter("id", id).RelatedSel("user").One(article)
	o.LoadRelated(article, "Comment", 2)
	return article
}
