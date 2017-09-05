package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

//DocSysvalue 系统变量表
// Label 在页面显示中的标签，或中文描述
// Creator 创建者
type DocSysvalue struct {
	Id int64 `orm:"column(id);pk;auto"`

	Kvgroup string `orm:"column(kvgroup);size(255)"`
	Key     string `orm:"column(key);size(255)"`

	Value      string `orm:"column(value);size(255)"`
	SortFactor uint64 `orm:"column(sort_factor);size(20);default(0)"`
	//中文描述
	Label string `orm:"column(label);size(255)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *DocSysvalue) TableName() string {
	return beego.AppConfig.String("doc_sysvalue_table")
}

func init() {
	orm.RegisterModel(new(DocSysvalue))
}

// AddDocSysvalue insert a new DocSysvalue into database and returns
// last inserted Id on success.
func AddDocSysvalue(m *DocSysvalue) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDocSysvalueById retrieves DocSysvalue by Id. Returns error if
// Id doesn't exist
func GetDocSysvalueById(id int64) (v *DocSysvalue, err error) {
	o := orm.NewOrm()
	v = &DocSysvalue{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDocSysvalue retrieves all DocSysvalue matches certain condition. Returns empty list if
// no records exist
func GetAllDocSysvalue(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DocSysvalue))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		qs = qs.Filter(k, v)
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []DocSysvalue
	qs = qs.OrderBy(sortFields...)
	if _, err := qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateDocSysvalue updates DocSysvalue by Id and returns error if
// the record to be updated doesn't exist
func UpdateDocSysvalueById(m *DocSysvalue) (err error) {
	o := orm.NewOrm()
	v := DocSysvalue{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDocSysvalue deletes DocSysvalue by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDocSysvalue(id int64) (err error) {
	o := orm.NewOrm()
	v := DocSysvalue{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DocSysvalue{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
