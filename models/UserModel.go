package models

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	. "github.com/devuser/beego-admin-demo/lib"
)

//用户表
type User struct {
	Id         int64  `orm:"column(id);pk;auto"`
	Username   string `orm:"column(username);unique;size(32)" form:"Username"  valid:"Required;MaxSize(32);MinSize(6)"`
	Password   string `orm:"column(password);size(32)" form:"Password" valid:"Required;MaxSize(20);MinSize(6)"`
	Repassword string `orm:"-" form:"Repassword" valid:"Required"`
	Nickname   string `orm:"size(32)" form:"Nickname" valid:"Required;MaxSize(20);MinSize(2)"`
	Realname   string `orm:"size(32)" form:"Realname" valid:"Required"`
	Title      string `orm:"size(32)" form:"Title" valid:"Required"`
	Delta      string `orm:"size(32)" form:"Delta" valid:"Required"`
	Phone      string `orm:"size(32)" form:"Phone" valid:"Required"`
	MobelPhone string `orm:"unique;size(32)" form:"MobilePhone" valid:"Required"`

	Email  string `orm:"unique;size(255)" form:"Email" valid:"Email"`
	Remark string `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`

	// Status int  `orm:"column(status);default(2)" form:"Status"`
	//锁定，对应删除动作
	Locked bool `orm:"column(locked);default(0);null" form:"Locked"`
	//用户上次修改密码的时间
	LastPasswordChangeAt time.Time `orm:"null;column(last_password_change_at)"`
	//用户首次登录的判断标准： last_login_time is null
	LastLoginTime time.Time `orm:"null;column(last_login_time)" form:"-"`
	//本次登录时间，2个小时内必须进行二次鉴权，2个小时由服务器设定
	LoginTime    time.Time `orm:"null;column(login_time)" form:"-"`
	LastLoginIP  string    `orm:"null;size(255);column(last_login_ip)" form:"-"`
	LoginIP      string    `orm:"null;size(255);column(login_ip)" form:"-"`
	LimitLoginIP string    `orm:"null;size(255);column(LIMIT_LOGIN_IP)" form:"-"`
	Online       bool      `orm:"default(0);null;column(is_online)" form:"-"`

	//错误密码次数，注意成功登录后则恢复到零
	LoginTryCount uint64 `orm:"column(login_try_count);default(0);size(20)" form:"LoginTryCount"`
	//最后一次尝试密码时间
	LoginTryLastAt time.Time `orm:"column(login_try_last_at);type(datetime);null"`
	//密码锁定
	LoginTryPasswordLocked bool `orm:"default(0);null;column(login_try_password_locked)" form:"LoginTryPasswordLocked"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`

	Role          []*Role      `orm:"rel(m2m)"`
	DocaccessList []*Docaccess `orm:"reverse(many)"`
	DocGroupList  []*DocGroup  `orm:"reverse(many)"`
	DocOrg        *DocOrg      `orm:"rel(fk)"`
}

func (u *User) TableName() string {
	return beego.AppConfig.String("rbac_user_table")
}

func (u *User) Valid(v *validation.Validation) {
	if u.Password != u.Repassword {
		v.SetError("Repassword", "两次输入的密码不一样")
	}
}

//验证用户信息
func checkUser(u *User) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&u)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

func init() {
	orm.RegisterModel(new(User))
}

/************************************************************/

//get user list
func Getuserlist(page int64, page_size int64, sort string) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	qs := o.QueryTable(user)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&users)
	count, _ = qs.Count()
	return users, count
}

//添加用户
func AddUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := new(User)
	user.Username = u.Username
	user.Password = Strtomd5(u.Password)
	user.Nickname = u.Nickname
	user.Email = u.Email
	user.Remark = u.Remark
	user.Status = u.Status

	id, err := o.Insert(user)
	return id, err
}

//更新用户
func UpdateUser(u *User) (int64, error) {
	if err := checkUser(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	user := make(orm.Params)
	if len(u.Username) > 0 {
		user["Username"] = u.Username
	}
	if len(u.Nickname) > 0 {
		user["Nickname"] = u.Nickname
	}
	if len(u.Email) > 0 {
		user["Email"] = u.Email
	}
	if len(u.Remark) > 0 {
		user["Remark"] = u.Remark
	}
	if len(u.Password) > 0 {
		user["Password"] = Strtomd5(u.Password)
	}
	if u.Status != 0 {
		user["Status"] = u.Status
	}
	if len(user) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table User
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(user)
	return num, err
}

func DelUserById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&User{Id: (Id)})
	return status, err
}

func GetUserByUsername(username string) (user User) {
	user = User{Username: username}
	o := orm.NewOrm()
	o.Read(&user, "username")
	fmt.Printf("User: %q\n", user)
	return user
}
