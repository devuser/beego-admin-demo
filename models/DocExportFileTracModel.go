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

//ExportFileTrac跟踪导出状态
//	如果不记录导出开始时间、导出结束时间的，那么可以统一设置为请求下载时间
type ExportFileTrac struct {
	Id             int64    `orm:"column(id);pk;auto"`
	Docname        *Docname `orm:"column(docname_id);rel(fk)"`
	ExportFileType string   `orm:"column(export_filetype);size(3)"`
	ExportFileName string   `orm:"column(export_filename);unique;size(255)"`

	ExportStartAt time.Time `orm:"column(export_start_at);type(datetime);null"`
	ExportOverAt  time.Time `orm:"column(export_over_at);type(datetime);null"`
	TransferAt    time.Time `orm:"column(transfer_at);type(datetime);null"`
	State         string    `orm:"column(state);size(4)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *ExportFileTrac) TableName() string {
	return beego.AppConfig.String("doc_exportfiletrac_table")
}

func init() {
	orm.RegisterModel(new(ExportFileTrac))
}

// AddExportFileTrac insert a new ExportFileTrac into database and returns
// last inserted Id on success.
func AddExportFileTrac(m *ExportFileTrac) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetExportFileTracById retrieves ExportFileTrac by Id. Returns error if
// Id doesn't exist
func GetExportFileTracById(id int64) (v *ExportFileTrac, err error) {
	o := orm.NewOrm()
	v = &ExportFileTrac{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllExportFileTrac retrieves all ExportFileTrac matches certain condition. Returns empty list if
// no records exist
func GetAllExportFileTrac(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(ExportFileTrac))
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

	var l []ExportFileTrac
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

// UpdateExportFileTrac updates ExportFileTrac by Id and returns error if
// the record to be updated doesn't exist
func UpdateExportFileTracById(m *ExportFileTrac) (err error) {
	o := orm.NewOrm()
	v := ExportFileTrac{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteExportFileTrac deletes ExportFileTrac by Id and returns error if
// the record to be deleted doesn't exist
func DeleteExportFileTrac(id int64) (err error) {
	o := orm.NewOrm()
	v := ExportFileTrac{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&ExportFileTrac{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
