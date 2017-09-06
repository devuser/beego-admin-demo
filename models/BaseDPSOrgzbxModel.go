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

// DPSOrgzbx 存款账户指标项
type DPSOrgzbx struct {
	Id int64 `orm:"column(id);pk;auto"`

	//归属机构
	Guishjig string `orm:"column(guishjig_code);size(32);null"`
	//服务机构
	Chjianjg string `orm:"column(chjianjg_code);size(32);null"`

	//开户日期
	Kaihriqi string `orm:"column(kaihriqi);size(8);"`
	//销户日期
	Xiohriqi string `orm:"column(xiohriqi);size(8);"`

	//存款类型_七种
	CunKuanLeiXing string `orm:"column(cunkuanleixing);size(4)"`

	//存款性质
	CunKuanXingZhi string `orm:"column(cunkuanxingzhi);size(255);null"`

	//客户并表标识
	KeHuBingBiaoFlag bool `orm:"column(kehubingbiao_flag);default(0);null"`
	//账户并表标识
	ZhangHuBingBiaoFlag bool `orm:"column(zhanghubingbiao_flag);default(0);null"`
	//存款产品Id
	CunKuanChanPinId int64 `orm:"column(cunkuanchanpin);size(20);null"`
	//期限
	QiXian int64 `orm:"column(qixian);size(20);null"`
	//期限单位Id
	QiXianUnitId int64 `orm:"column(qixianunit_id);size(20);default(3)"`
	//时点余额
	ShiDianYE float64 `orm:"column(shidianye);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//币种编码
	Huobdaih string `orm:"column(huobdhao);type(string);size(2);default(01)"  form:"Huobdaih"  valid:"Required"`

	// BizdateId
	Bizdate string `orm:"column(bizdate);size(8)"`
	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

// TableName 缺省方法，返回对应的数据库表名
func (t *DPSOrgzbx) TableName() string {
	return beego.AppConfig.String("base_dps_orgckzhzbx_table")
}

func init() {
	orm.RegisterModel(new(DPSOrgzbx))
}

// AddDPSOrgzbx insert a new DPSOrgzbx into database and returns
// last inserted Id on success.
func AddDPSOrgzbx(m *DPSOrgzbx) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDPSOrgzbxById retrieves DPSOrgzbx by Id. Returns error if
// Id doesn't exist
func GetDPSOrgzbxById(id int64) (v *DPSOrgzbx, err error) {
	o := orm.NewOrm()
	v = &DPSOrgzbx{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetDPSOrgzbxList(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(DPSOrgzbx)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}

	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id",
		"Guishjig",
		"Chjianjg",
		"Kaihriqi",
		"Xiohriqi",
		"CunKuanLeiXing",
		"CunKuanXingZhi",
		"KeHuBingBiaoFlag",
		"ZhangHuBingBiaoFlag",
		"CunKuanChanPinId",
		"Status")
	count, _ = qs.Count()
	return nodes, count
}

// GetAllDPSOrgzbx retrieves all DPSOrgzbx matches certain condition. Returns empty list if
// no records exist
func GetAllDPSOrgzbxs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DPSOrgzbx))
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

	var l []DPSOrgzbx
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

// UpdateDPSOrgzbx updates DPSOrgzbx by Id and returns error if
// the record to be updated doesn't exist
func UpdateDPSOrgzbxById(m *DPSOrgzbx) (err error) {
	o := orm.NewOrm()
	v := DPSOrgzbx{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDPSOrgzbx deletes DPSOrgzbx by Id and returns error if
// the record to be deleted doesn't exist
func DelDPSOrgzbxById(id int64) (err error) {
	o := orm.NewOrm()
	v := DPSOrgzbx{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DPSOrgzbx{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
