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

//DocOrg 基础组织机构树
//来自核心系统
//
// Orgcode 错格式的编码格式
// StdOrgcode Gopher规定的组织机构编码
type DocOrg struct {
	Id          int64  `orm:"column(id);pk;auto"`
	Orgname     string `orm:"column(orgname);size(255)"`
	FullOrgname string `orm:"column(full_orgname);unique;size(255)"`

	Orgcode    string `orm:"column(orgcode);unique;size(6);null"`
	StdOrgcode string `orm:"column(stdorgcode);unique;size(6);null"`

	//
	ParentId    int64 `orm:"column(parent_id);size(20);null"`
	StdParentId int64 `orm:"column(std_parent_id);size(20);null"`

	// 设置一对多的反向关系
	UserList []*User `orm:"reverse(many)"`
	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *DocOrg) TableName() string {
	return beego.AppConfig.String("doc_org_table")
}

func init() {
	orm.RegisterModel(new(DocOrg))
}

// AddDocOrg insert a new DocOrg into database and returns
// last inserted Id on success.
func AddDocOrg(m *DocOrg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDocOrgById retrieves DocOrg by Id. Returns error if
// Id doesn't exist
func GetDocOrgById(id int64) (v *DocOrg, err error) {
	o := orm.NewOrm()
	v = &DocOrg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDocOrg retrieves all DocOrg matches certain condition. Returns empty list if
// no records exist
func GetAllDocOrg(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DocOrg))
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

	var l []DocOrg
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

// UpdateDocOrg updates DocOrg by Id and returns error if
// the record to be updated doesn't exist
func UpdateDocOrgById(m *DocOrg) (err error) {
	o := orm.NewOrm()
	v := DocOrg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDocOrg deletes DocOrg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDocOrg(id int64) (err error) {
	o := orm.NewOrm()
	v := DocOrg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DocOrg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
