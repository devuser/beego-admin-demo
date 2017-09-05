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

//DocDataorg 数据权限组织树
//
//	暂时没有使用
//	一般情况下与组织机构树相同，当然也可以不同
//	一个用户可以与数据权限组织树上的多个组织对应，但是多个组织不可以有交叉
type DocDataorg struct {
	Id          int64  `orm:"column(id);pk;auto"`
	Orgname     string `orm:"column(orgname);size(255)"`
	FullOrgname string `orm:"column(full_orgname);unique;size(255)"`
	Orgcode     string `orm:"column(orgcode);unique;size(6);null"`
	StdOrgcode  string `orm:"column(stdorgcode);unique;size(6);null"`

	ParentId    int64 `orm:"column(parent_id);size(20);null"`
	StdParentId int64 `orm:"column(std_parent_id);size(20);null"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *DocDataorg) TableName() string {
	return beego.AppConfig.String("doc_dataorg_table")
}

func init() {
	orm.RegisterModel(new(DocDataorg))
}

// AddDocDataorg insert a new DocDataorg into database and returns
// last inserted Id on success.
func AddDocDataorg(m *DocDataorg) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDocDataorgById retrieves DocDataorg by Id. Returns error if
// Id doesn't exist
func GetDocDataorgById(id int64) (v *DocDataorg, err error) {
	o := orm.NewOrm()
	v = &DocDataorg{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDocDataorg retrieves all DocDataorg matches certain condition. Returns empty list if
// no records exist
func GetAllDocDataorg(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DocDataorg))
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

	var l []DocDataorg
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

// UpdateDocDataorg updates DocDataorg by Id and returns error if
// the record to be updated doesn't exist
func UpdateDocDataorgById(m *DocDataorg) (err error) {
	o := orm.NewOrm()
	v := DocDataorg{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDocDataorg deletes DocDataorg by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDocDataorg(id int64) (err error) {
	o := orm.NewOrm()
	v := DocDataorg{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DocDataorg{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
