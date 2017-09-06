package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"reflect"
	"strings"
	"time"
)

//BaseCurrency 数据权限组织树
//一般情况下与组织机构树相同，当然也可以不同
//一个用户可以与数据权限组织树上的多个组织对应，但是多个组织不可以有交叉
type BaseCurrency struct {
	Id int64 `orm:"column(id);pk;auto"`
	//货币代号
	Huobdaih string `orm:"column(huobdaih);unique;size(2);valid(required)"`
	//币种别名
	Bzbiemng string `orm:"column(bzbiemng);unique;size(32);valid(required)"`
	//货币名称
	Hbmcheng string `orm:"column(hbmcheng);unique;size(32);valid(required)"`
	//货币字母缩写
	Huobzmsx string `orm:"column(huobzmsx);size(32);valid(required)"`
	//货币符号
	Huobfhao string `orm:"column(huobfhao);size(32);valid(required)"`
	//国别代码
	Guobiedm string `orm:"column(guobiedm);size(32);valid(required)"`
	//辅币进位
	Fubijinw string `orm:"column(fubijinw);size(32);valid(required)"`
	//牌价最低位数
	Pjzdweis string `orm:"column(pjzdweis);size(32);valid(required)"`
	//最小货币单位
	Zuixhbdw string `orm:"column(zuixhbdw);size(32);valid(required)"`
	//最小计息单位
	Zuixjxdw string `orm:"column(zuixjxdw);size(32);valid(required)"`
	//最小记帐单位
	Zuixjzdw string `orm:"column(zuixjzdw);size(32);valid(required)"`
	//折角分位
	Zjfenwei string `orm:"column(zjfenwei);size(32);valid(required)"`
	//现金取现限额
	Xjinqxxe string `orm:"column(xjinqxxe);size(32);valid(required)"`
	//货币等级
	Huobidji string `orm:"column(huobidji);size(32);valid(required)"`
	//加入其它货币标志
	Jiarubzz string `orm:"column(jiarubzz);size(32);valid(required)"`
	//加入货币代号
	Bizhongg string `orm:"column(bizhongg);size(32);valid(required)"`
	//英联邦货币标志
	Yinglbhb string `orm:"column(yinglbhb);size(32);valid(required)"`
	//说明
	Shenming string `orm:"column(shenming);size(32);valid(required)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

// TableName 缺省方法，返回对应的数据库表名
func (t *BaseCurrency) TableName() string {
	return beego.AppConfig.String("base_common_currency_table")
}

func init() {
	orm.RegisterModel(new(BaseCurrency))
}

// AddBaseCurrency insert a new BaseCurrency into database and returns
// last inserted Id on success.
func AddBaseCurrency(m *BaseCurrency) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBaseCurrencyById retrieves BaseCurrency by Id. Returns error if
// Id doesn't exist
func GetBaseCurrencyById(id int64) (v *BaseCurrency, err error) {
	o := orm.NewOrm()
	v = &BaseCurrency{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//get node list
func GetBaseCurrencylist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseCurrency)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}

	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id",
		"huobdaih",
		"Bzbiemng",
		"Hbmcheng",
		"Huobzmsx",
		"Huobfhao",
		"Guobiedm",
		"Fubijinw",
		"Pjzdweis",
		"Zuixhbdw",
		"Zuixjxdw",
		"Zuixjzdw",
		"Zjfenwei",
		"Xjinqxxe",
		"Huobidji",
		"Jiarubzz",
		"Bizhongg",
		"Yinglbhb",
		"Shenming",
		"Status")
	count, _ = qs.Count()
	return nodes, count
}

// GetAllBaseCurrency retrieves all BaseCurrency matches certain condition. Returns empty list if
// no records exist
func GetAllBaseCurrency(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BaseCurrency))
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

	var l []BaseCurrency
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

//验证用户信息
func checkBaseCurrency(u *BaseCurrency) (err error) {
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

func UpdateBaseCurrency(n *BaseCurrency) (int64, error) {
	if err := checkBaseCurrency(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateBaseCurrency.Huobdaih %s\n", n.Huobdaih)
	o := orm.NewOrm()
	node := make(orm.Params)
	if len(n.Huobdaih) > 0 {
		node["Huobdaih"] = n.Huobdaih
	}
	// if len(n.Name) > 0 {
	// 	node["Name"] = n.Name
	// }
	// if len(n.Remark) > 0 {
	// 	node["Remark"] = n.Remark
	// }
	// if n.Status != 0 {
	// 	node["Status"] = n.Status
	// }
	if len(node) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table BaseCurrency
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

// DeleteBaseCurrency deletes BaseCurrency by Id and returns error if
// the record to be deleted doesn't exist

func DelBaseCurrencyById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&BaseCurrency{Id: Id})
	return status, err
}
