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

//BaseSTMQudaoTz 渠道-企业转账台账
//
// 	通过核心系统交易流水取数：
// 	【Gopher交易来源为“电财通”】 且 【账户大类为“客户账户”】的流水明细交易金额合计/流水明细交易记录个数
// 	5．开通电财通电子支付账户数量、客户数量
// 	通过电财通系统内管端取数：取“资金管理——转账汇款”类别下开通“跨行转账”功能的账户数量、客户数量。
type BaseSTMQudaoTz struct {
	Id int64 `orm:"column(id);pk;auto"`
	//业务日期
	Bizdate string `orm:"column(bizdate);size(8)"`
	//开户机构（报表显示为填报机构）
	Orgcode string `orm:"column(org_code);size(32);null"`
	//通过电财通电子支付结算资金量
	DctZjzl float64 `orm:"column(dct_zjzl);size(20);default(0)"`
	//通过电财通电子支付结算笔数
	DctJsbs int64 `orm:"column(dct_jsbs);size(20);default(0)"`
	//开通电财通电子支付账户数量
	DctAccountCount int64 `orm:"column(dct_account_count);size(20);default(0)"`
	//开通电财通电子支付客户数量
	DctCustomCount int64 `orm:"column(dct_custom_count);size(20);default(0)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

// //节点表
// type BaseSTMQudaoTz struct {
// 	SQLModel `sql:",inline"`
//
// 	Bizdate time.Time `orm:"unique;column(bizdate);type(datetime);null"`
// 	//@todo TODO DONE
// 	State string `orm:"unique;column(bizdate);type(datetime);null"	form:"State"  valid:"Required"`
// }

func (n *BaseSTMQudaoTz) TableName() string {
	return beego.AppConfig.String("base_stmqudaotz_table")
}

//验证用户信息
func checkBaseSTMQudaoTz(u *BaseSTMQudaoTz) (err error) {
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
	orm.RegisterModel(new(BaseSTMQudaoTz))
}

//get node list
func GetBaseSTMQudaoTzlist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTz)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "Bizdate", "State", "Status")
	count, _ = qs.Count()
	return nodes, count
}

func GetBaseSTMQudaoTzById(nid int64) (BaseSTMQudaoTz, error) {
	o := orm.NewOrm()
	node := BaseSTMQudaoTz{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddBaseSTMQudaoTz(n *BaseSTMQudaoTz) (int64, error) {
	if err := checkBaseSTMQudaoTz(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTz)
	// node.Name = n.Name
	// node.State = n.State
	// node.Country = n.Country

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateBaseSTMQudaoTz(n *BaseSTMQudaoTz) (int64, error) {
	if err := checkBaseSTMQudaoTz(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := make(orm.Params)
	//@todo 校验条件
	// if len(node) == 0 {
	// 	return 0, errors.New("update field is empty")
	// }
	var table BaseSTMQudaoTz
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelBaseSTMQudaoTzById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&BaseSTMQudaoTz{Id: Id})
	return status, err
}

func GetBaseSTMQudaoTzlistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTz)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetBaseSTMQudaoTzTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTz)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllDownloadtrac retrieves all Downloadtrac matches certain condition. Returns empty list if
// no records exist
func GetAllBaseSTMQudaoTzs(query map[string]string, fields []string, sortby []string, order []string,
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
