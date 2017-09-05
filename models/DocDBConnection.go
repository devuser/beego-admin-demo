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

//DocDBConnection 系统变量表
// Label 在页面显示中的标签，或中文描述
// Creator 创建者
type DocDBConnection struct {
	Id      int64     `orm:"column(id);pk;auto"`
	StartAt time.Time `orm:"column(start_at);type(datetime)"`
	Used    bool      `orm:"column(used);"`

	DataSourceName string `orm:"column(datasource_name);unique;size(255)"`

	SpringDataSourceURL string `orm:"column(url);size(255);valid(required)"`

	SpringDataSourceUsername                  string `orm:"column(username);size(255);valid(required)"`
	SpringDataSourcePassword                  string `orm:"column(password);size(255);valid(required)"`
	SpringDataSourceDriverClassName           string `orm:"column(driver_class_name);size(255);valid(required)"`
	SpringDataSourceSchema                    string `orm:"column(schema);size(255);valid(required)"`
	SpringDataSourceMaxWait                   uint64 `orm:"column(max_wait);size(20)"`
	SpringDataSourceMaxActive                 uint64 `orm:"column(max_active);range(0,10000)"`
	SpringDataSourcePoolName                  string `orm:"column(pool_name);size(255);default(hikariCP)"`
	SpringDataSourceMaximumPoolSize           uint64 `orm:"column(pool_size);range(0,10000)"`
	SpringDataSourceMinimumIdle               uint64 `orm:"column(minimum_idle);range(0,10000)"`
	SpringDataSourceConnectionTimeout         uint64 `orm:"column(connection_timeout)"`
	SpringDataSourceIdleTimeout               uint64 `orm:"column(idle_timeout)"`
	SpringDataSourcePoolPreparedStatements    bool   `orm:"column(pool_prepared_statements)"`
	SpringDataSourceMaxOpenPreparedStatements uint64 `orm:"column(max_open_prepared_statements)"`
	//管控部分
	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *DocDBConnection) TableName() string {
	return beego.AppConfig.String("doc_dbconnection_table")
}

func init() {
	orm.RegisterModel(new(DocDBConnection))
}

// AddDocDBConnection insert a new DocDBConnection into database and returns
// last inserted Id on success.
func AddDocDBConnection(m *DocDBConnection) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDocDBConnectionById retrieves DocDBConnection by Id. Returns error if
// Id doesn't exist
func GetDocDBConnectionById(id int64) (v *DocDBConnection, err error) {
	o := orm.NewOrm()
	v = &DocDBConnection{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllDocDBConnection retrieves all DocDBConnection matches certain condition. Returns empty list if
// no records exist
func GetAllDocDBConnection(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DocDBConnection))
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

	var l []DocDBConnection
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

// UpdateDocDBConnection updates DocDBConnection by Id and returns error if
// the record to be updated doesn't exist
func UpdateDocDBConnectionById(m *DocDBConnection) (err error) {
	o := orm.NewOrm()
	v := DocDBConnection{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDocDBConnection deletes DocDBConnection by Id and returns error if
// the record to be deleted doesn't exist
func DeleteDocDBConnection(id int64) (err error) {
	o := orm.NewOrm()
	v := DocDBConnection{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DocDBConnection{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
