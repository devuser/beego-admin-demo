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

//BIZDATE, SAP_ORG_CODE, ACCOUNTTITLE_CODE, COIN_TYPE_CODE, ACCOUNT_SETS, BALANCE, STATUS, CREATOR, CREATE_AT, UPDATER, UPDATE_AT, MESSAGE

// SAPDayBal SAP日终余额表
type SAPDayBal struct {
	Id int64 `orm:"column(id);auto;pk"`

	//业务日期
	Bizdate string `orm:"column(BIZDATE);size(8);"`
	//SAP组织机构编码
	SAPOrgcode string `orm:"column(SAP_ORG_CODE);size(32);"`
	//科目代号
	AccountTitleCode string `orm:"column(ACCOUNTTITLE_CODE);size(10);"`
	//币种
	CoinTypeCode string `orm:"column(COIN_TYPE_CODE);size(10);"`
	//帐套 营业：0  财务：1合计：2
	AccountSets string `orm:"column(ACCOUNT_SETS);size(1);"`
	//余额
	Balance float64 `orm:"column(BALANCE);type(float64-decimal);decimals(2);digits(23);default(0)"`

	//管控部分
	Status   uint64    `orm:"column(STATUS);default(2)" form:"Status"`
	Creator  *User     `orm:"column(CREATOR);null;rel(fk)"`
	CreateAt time.Time `orm:"column(CREATE_AT);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(UPDATER);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(UPDATE_AT);type(datetime);null"`
	Message  string    `orm:"column(MESSAGE);size(255);null"`
}

func (n *SAPDayBal) TableName() string {
	return beego.AppConfig.String("base_sapdaybal_table")
}

//验证用户信息
func checkSAPDayBal(u *SAPDayBal) (err error) {
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
	orm.RegisterModel(new(SAPDayBal))
}

//get node list
func GetSAPDayBallist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(SAPDayBal)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "SAPDayBal", "Description", "Cabinet", "Folder", "FileType", "Paramdesc", "ETLTaskDescId", "UpdateAt", "CreateAt")
	count, _ = qs.Count()
	return nodes, count
}

func GetSAPDayBalById(nid int64) (SAPDayBal, error) {
	o := orm.NewOrm()
	node := SAPDayBal{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddSAPDayBal(n *SAPDayBal) (int64, error) {
	if err := checkSAPDayBal(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(SAPDayBal)
	// node.DianPiaoCode = n.DianPiaoCode
	// node.DocOrg = n.DocOrg
	// node.BizdateId = n.BizdateId

	// node.DebitOrCredit = n.DebitOrCredit
	// node.ProductType = n.ProductType

	node.Balance = n.Balance
	// node.Interest = n.Interest
	// node.InsidePlanFlag = n.InsidePlanFlag
	// node.StartAt = n.StartAt
	// node.OverAt = n.OverAt

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateSAPDayBal(n *SAPDayBal) (int64, error) {
	if err := checkSAPDayBal(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateSAPDayBal.SAPDayBal %q\n", n)
	o := orm.NewOrm()
	node := make(orm.Params)
	// if len(n.DianPiaoCode) > 0 {
	// 	node["DianPiaoCode"] = n.DianPiaoCode
	// }
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
	var table SAPDayBal
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelSAPDayBalById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&SAPDayBal{Id: Id})
	return status, err
}

func GetSAPDayBallistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(SAPDayBal)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetSAPDayBalTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(SAPDayBal)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllSAPDayBal retrieves all SAPDayBal matches certain condition. Returns empty list if
// no records exist
func GetAllSAPDayBals(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SAPDayBal))
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

	var l []SAPDayBal
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
