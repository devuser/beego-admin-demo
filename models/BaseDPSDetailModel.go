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

// DPSDetail 存款账户日终表
type DPSDetail struct {
	Id int64 `orm:"column(id);pk;auto"`

	//账户号
	ZhangHao string `orm:"column(zhanghao);size(64);null"`
	//账户名称
	Zhuzwmc string `orm:"column(zhhuzwmc);size(255);null"`
	//子账号（序号）
	Zhhaoxuh string `orm:"column(zhhaoxuh);size(32);null"`

	//借贷标志
	//0--借	借-余额减少贷-余额增加
	//1--贷	借-余额减少贷-余额增加
	Jiedaibz bool `orm:"column(jiedaibz);default(0);"`
	//交易币种
	Jiaoybiz string `orm:"column(jiaoybiz);type(string);size(2);default(01)"  form:"jiaoybiz"  valid:"Required"`

	//JIAOYIJE 交易金额
	Jiaoyije float64 `orm:"column(jiaoyije);type(float64-decimal);decimals(2);digits(16);null"`
	//ZHANGHYE 账户余额
	Zhanghye float64 `orm:"column(zhanghye);type(float64-decimal);decimals(2);digits(23);null"`

	//累计利息
	Interest float64 `orm:"column(interest);type(float64-decimal);decimals(2);digits(23);null"`

	//QDAOLEIX 交易渠道
	Qdaoleix string `orm:"column(qdaoleix);type(string);size(3);"  form:"QDAOLEIX"  valid:"Required"`

	//交易日期
	Jiaoyirq string `orm:"column(jiaoyirq);size(8);"`
	//主机日期
	Zhujriqi string `orm:"column(zhujriqi);size(8);"`

	//CHONGZBZ	冲正标志	varchar2(1)	N
	// 0--无关
	// 1--当日冲正
	// 2--隔日冲正

	Chongzbz string `orm:"column(chongzbz);default(0);size(1);null"`
	Bchongbz string `orm:"column(bchongbz);default(0);size(1);null"`
	//CUOZRIQI	错账原日期	varchar2(8)	Y
	Cuozriqi string `orm:"unique;column(cuozriqi);size(8);"`
	//CUOZLIUS	错账原柜员流水号	varchar2(32)	Y
	Cuozlius string `orm:"column(cuozlius);size(64);null"`

	//============= 与标签有有关结束 =============
	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

// TableName 缺省方法，返回对应的数据库表名
func (t *DPSDetail) TableName() string {
	return beego.AppConfig.String("base_dps_detail_table")
}

func init() {
	orm.RegisterModel(new(DPSDetail))
}

// AddDPSDetail insert a new DPSDetail into database and returns
// last inserted Id on success.
func AddDPSDetail(m *DPSDetail) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDPSDetailById retrieves DPSDetail by Id. Returns error if
// Id doesn't exist
func GetDPSDetailById(id int64) (v *DPSDetail, err error) {
	o := orm.NewOrm()
	v = &DPSDetail{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetDPSDetailList(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(DPSDetail)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "ZhangHao", "Zhuzwmc", "Zhhaoxuh", "Status")
	count, _ = qs.Count()
	return nodes, count
}

// GetAllDPSDetail retrieves all DPSDetail matches certain condition. Returns empty list if
// no records exist
func GetAllDPSDetails(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DPSDetail))
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

	var l []DPSDetail
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

// UpdateDPSDetail updates DPSDetail by Id and returns error if
// the record to be updated doesn't exist
func UpdateDPSDetailById(m *DPSDetail) (err error) {
	o := orm.NewOrm()
	v := DPSDetail{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDPSDetail deletes DPSDetail by Id and returns error if
// the record to be deleted doesn't exist
func DelDPSDetailById(id int64) (err error) {
	o := orm.NewOrm()
	v := DPSDetail{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DPSDetail{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
