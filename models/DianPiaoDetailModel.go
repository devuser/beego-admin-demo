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

// Name	Code	Data Type	Length	Precision	Primary	Foreign Key	Mandatory
// 电票交易编号	DianPiaoCode	varchar(32)	32		FALSE	FALSE	FALSE
// 组织机构Id	OrgId	bigint(20)	20		FALSE	FALSE	FALSE
// BizdateId	BizdateId	bigint(20)	20		FALSE	FALSE	FALSE
// 发放或回收	debit_or_credit	char(4)	4		FALSE	FALSE	FALSE
// 产品类型	产品类型	bigint(20)	20		FALSE	FALSE	FALSE
// 余额	余额	decimal(19,2)	19	2	FALSE	FALSE	FALSE
// 贴现利息	贴现利息	decimal(19,2)	19	2	FALSE	FALSE	FALSE
// 是否计划内	InsidePlanFlag	boolean			FALSE	FALSE	FALSE
// 生效日期	生效日期	char(8)	8		FALSE	FALSE	FALSE
// 失效日期	失效日期	char(8)	8		FALSE	FALSE	FALSE
// Creator	Creator	varchar(32)	32		FALSE	FALSE	FALSE
// CreateAt	create_at	datetime			FALSE	FALSE	FALSE

// DianPiaoDetails 电票交易明细
type DianPiaoDetails struct {
	Id           int64  `orm:"column(id);auto;pk"`
	DianPiaoCode string `orm:"column(dianpiao_code);size(32);unique" form:"dianpiao_code" valid:"Required"`
	Orgcode      string `orm:"column(org_code);size(32);null"`

	DebitOrCredit string `orm:"column(debit_or_credit);size(4)" form:"debit_or_credit"  valid:"Required"`
	//产品类型
	ProductType *DianPiaoProductType `orm:"column(producttype);rel(fk)"`
	//余额和利息
	Balance  float64 `orm:"column(Balance);type(float64-decimal);decimals(2);digits(23);default(0)"`
	Interest float64 `orm:"column(Interest);type(float64-decimal);decimals(2);digits(23);default(0)"`

	//计划内
	InsidePlanFlag bool `orm:"column(InsidePlanFlag);default(0)""`

	//生效日期、失效日期
	StartAt string `orm:"column(start_at);size(8)"`
	OverAt  string `orm:"column(over_at);size(8)"`

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

func (n *DianPiaoDetails) TableName() string {
	return beego.AppConfig.String("base_dianpiaodetails_table")
}

//验证用户信息
func checkDianPiaoDetails(u *DianPiaoDetails) (err error) {
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
	orm.RegisterModel(new(DianPiaoDetails))
}

//get node list
func GetDianPiaoDetailslist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(DianPiaoDetails)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "DianPiaoDetails", "Description", "Cabinet", "Folder", "FileType", "Paramdesc", "ETLTaskDescId", "UpdateAt", "CreateAt")
	count, _ = qs.Count()
	return nodes, count
}

func GetDianPiaoDetailsById(nid int64) (DianPiaoDetails, error) {
	o := orm.NewOrm()
	node := DianPiaoDetails{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

const (
	cstFirstOrgcode = "000000"
)

//添加用户
func AddDianPiaoDetails(n *DianPiaoDetails) (int64, error) {
	if err := checkDianPiaoDetails(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(DianPiaoDetails)
	node.DianPiaoCode = n.DianPiaoCode
	node.Orgcode = cstFirstOrgcode
	node.Bizdate = time.Now().Format("2006-01-02")

	node.DebitOrCredit = n.DebitOrCredit
	node.ProductType = n.ProductType

	node.Balance = n.Balance
	node.Interest = n.Interest
	node.InsidePlanFlag = n.InsidePlanFlag
	node.StartAt = n.StartAt
	node.OverAt = n.OverAt

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateDianPiaoDetails(n *DianPiaoDetails) (int64, error) {
	if err := checkDianPiaoDetails(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateDianPiaoDetails.DianPiaoDetails %q\n", n)
	o := orm.NewOrm()
	node := make(orm.Params)
	if len(n.DianPiaoCode) > 0 {
		node["DianPiaoCode"] = n.DianPiaoCode
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
	var table DianPiaoDetails
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelDianPiaoDetailsById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&DianPiaoDetails{Id: Id})
	return status, err
}

func GetDianPiaoDetailslistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(DianPiaoDetails)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetDianPiaoDetailsTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(DianPiaoDetails)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllDianPiaoDetails retrieves all DianPiaoDetails matches certain condition. Returns empty list if
// no records exist
func GetAllDianPiaoDetailss(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DianPiaoDetails))
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

	var l []DianPiaoDetails
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
