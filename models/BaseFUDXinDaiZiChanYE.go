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

// FUDXinDaiZiChanYE base_信贷资产余额表
type FUDXinDaiZiChanYE struct {
	Id int64 `orm:"column(id);pk;auto"`

	Orgcode string `orm:"column(org_code);size(32);null"`
	//货币代号
	Huobdaih string `orm:"column(huobdaih);type(string);size(2);default(01)"`

	//信贷资产余额
	XinDaiZiChanYE float64 `orm:"column(xindaizichanye);type(float64-decimal);decimals(2);digits(23);default(0)"`

	//业务日期
	Bizdate string `orm:"column(bizdate);size(8)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *FUDXinDaiZiChanYE) TableName() string {
	return beego.AppConfig.String("base_fud_xdzcye_table")
}

func init() {
	orm.RegisterModel(new(FUDXinDaiZiChanYE))
}

// AddFUDXinDaiZiChanYE insert a new FUDXinDaiZiChanYE into database and returns
// last inserted Id on success.
func AddFUDXinDaiZiChanYE(m *FUDXinDaiZiChanYE) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetFUDXinDaiZiChanYEById retrieves FUDXinDaiZiChanYE by Id. Returns error if
// Id doesn't exist
func GetFUDXinDaiZiChanYEById(id int64) (v *FUDXinDaiZiChanYE, err error) {
	o := orm.NewOrm()
	v = &FUDXinDaiZiChanYE{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllFUDXinDaiZiChanYE retrieves all FUDXinDaiZiChanYE matches certain condition. Returns empty list if
// no records exist
func GetAllFUDXinDaiZiChanYE(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FUDXinDaiZiChanYE))
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

	var l []FUDXinDaiZiChanYE
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

// UpdateFUDXinDaiZiChanYE updates FUDXinDaiZiChanYE by Id and returns error if
// the record to be updated doesn't exist
func UpdateFUDXinDaiZiChanYEById(m *FUDXinDaiZiChanYE) (err error) {
	o := orm.NewOrm()
	v := FUDXinDaiZiChanYE{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteFUDXinDaiZiChanYE deletes FUDXinDaiZiChanYE by Id and returns error if
// the record to be deleted doesn't exist
func DeleteFUDXinDaiZiChanYE(id int64) (err error) {
	o := orm.NewOrm()
	v := FUDXinDaiZiChanYE{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&FUDXinDaiZiChanYE{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
