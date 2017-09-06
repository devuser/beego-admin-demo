package models

import (
	"errors"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"reflect"
	"strings"
	"time"
)

//BaseDateUnit 周期单位，与长亮保持一致
type BaseDateUnit struct {
	Id int64 `orm:"column(id);pk;auto"`
	//周期编码
	DateUnit string `orm:"unique;column(dateunit);size(32)"`
	//周期描述
	Description string `orm:"column(description);size(255)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

// TableName 缺省方法，返回对应的数据库表名
func (n *BaseDateUnit) TableName() string {
	return beego.AppConfig.String("base_common_dateunit_table")
}

//验证用户信息
func checkBaseDateUnit(u *BaseDateUnit) (err error) {
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
	orm.RegisterModel(new(BaseDateUnit))
}

//get node list
func GetBaseDateUnitlist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseDateUnit)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "Title", "Name", "Status", "Pid", "Remark", "Group__id")
	count, _ = qs.Count()
	return nodes, count
}

func GetBaseDateUnitById(nid int64) (BaseDateUnit, error) {
	o := orm.NewOrm()
	node := BaseDateUnit{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddBaseDateUnit(n *BaseDateUnit) (int64, error) {
	if err := checkBaseDateUnit(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(BaseDateUnit)
	// node.Name = n.Name
	// node.State = n.State
	// node.Country = n.Country

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateBaseDateUnitById(n *BaseDateUnit) (int64, error) {
	if err := checkBaseDateUnit(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := make(orm.Params)
	//@todo 校验条件
	// if len(node) == 0 {
	// 	return 0, errors.New("update field is empty")
	// }
	var table BaseDateUnit
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelBaseDateUnitById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&BaseDateUnit{Id: Id})
	return status, err
}

func GetBaseDateUnitlistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseDateUnit)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetBaseDateUnitTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(BaseDateUnit)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllDownloadtrac retrieves all Downloadtrac matches certain condition. Returns empty list if
// no records exist
func GetAllBaseDateUnit(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Downloadtrac))
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

	var l []Downloadtrac
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
