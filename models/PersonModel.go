package models

import (
	"errors"
	"log"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	// . "github.com/devuser/beego-admin-demo/lib"
	"time"
)

//用户表
type Person struct {
	Id        int64  `orm:"column(id);pk;auto"`
	Firstname string `orm:"size(255)" form:"first_name"`
	Lastname  string `orm:"size(255)" form:"last_name"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (u *Person) TableName() string {
	return beego.AppConfig.String("sample_person_table")
}

func (u *Person) Valid(v *validation.Validation) {
	// if u.Password != u.Repassword {
	// 	v.SetError("Repassword", "两次输入的密码不一样")
	// }
}

//验证用户信息
func checkPerson(u *Person) (err error) {
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
	orm.RegisterModel(new(Person))
}

/************************************************************/

//get person list
func Getpersonlist(page int64, page_size int64, sort string) (persons []orm.Params, count int64) {
	o := orm.NewOrm()
	person := new(Person)
	qs := o.QueryTable(person)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&persons)
	count, _ = qs.Count()
	return persons, count
}

//添加用户
func AddPerson(u *Person) (int64, error) {
	if err := checkPerson(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	person := new(Person)
	person.Firstname = u.Firstname
	person.Lastname = u.Lastname
	//Strtomd5(u.Password)

	id, err := o.Insert(person)
	return id, err
}

//更新用户
func UpdatePerson(u *Person) (int64, error) {
	if err := checkPerson(u); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	person := make(orm.Params)
	if len(u.Firstname) > 0 {
		person["Firstname"] = u.Firstname
	}
	if len(u.Lastname) > 0 {
		person["Lastname"] = u.Lastname
	}

	if len(person) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Person
	num, err := o.QueryTable(table).Filter("Id", u.Id).Update(person)
	return num, err
}

func DelPersonById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Person{Id: (Id)})
	return status, err
}

// func GetPersonByPersonname(personname string) (person Person) {
// 	person = Person{Personname: personname}
// 	o := orm.NewOrm()
// 	o.Read(&person, "Personname")
// 	return person
// }
