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

// FBusSubCM 国内保理-关联合同表
type FBusSubCM struct {
	Id int64 `orm:"column(id);auto;pk"`

	Contractid      float64   `orm:"type(float64-decimal);default(0)"`
	Sellerid        string    `orm:"size(32)"`
	Buyerid         string    `orm:"size(32)"`
	Subctype        string    `orm:"size(1)"`
	Cofactor        string    `orm:"size(32)"`
	Insurance       string    `orm:"size(32)"`
	Thirdpart       string    `orm:"size(32)"`
	Zkstartdate     time.Time `orm:"type(datetime);null"`
	repokxq         float64   `orm:"type(float64-decimal);default(0)"`
	Dbfktimezone    float64   `orm:"type(float64-decimal);default(0)"`
	Loanpcnt        float64   `orm:"type(float64-decimal);decimals(8);digits(39);default(0)"`
	Coverpcnt       float64   `orm:"type(float64-decimal);decimals(8);digits(39);default(0)"`
	Ifcoverpcnt     float64   `orm:"type(float64-decimal);decimals(8);digits(39);default(0)"`
	Loankxq         float64   `orm:"type(float64-decimal);default(0)"`
	Transferway     string    `orm:"size(1)"`
	Interincomeflag string    `orm:"size(1)"`
	Intercloseflag  string    `orm:"size(1)"`
	Interpayflag    string    `orm:"size(1)"`
	Cycleflag       string    `orm:"size(1)"`
	Feflag          string    `orm:"size(10)"`
	rateflag        string    `orm:"size(10)"`
	Inactflag       string    `orm:"size(1)"`
	Inactdate       time.Time `orm:"type(datetime);null"`
	Clstartdate     time.Time `orm:"type(datetime);null"`
	Clenddate       time.Time `orm:"type(datetime);null"`
	Opstate         float64   `orm:"type(float64-decimal);default(0)"`
	WorkItemId      float64   `orm:"type(float64-decimal);default(0)"`
	State           float64   `orm:"type(float64-decimal);default(0)"`
	Dbprototal      float64   `orm:"column(dbprototal);type(float64-decimal);decimals(8);digits(39);default(0)"`
	Sellernr        string    `orm:"size(35)"`
	Buyernr         string    `orm:"size(35)"`

	//@TODO: 业务日期
	Bizdate string `orm:"column(bizdate);size(8)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (n *FBusSubCM) TableName() string {
	return beego.AppConfig.String("base_fbussubcm_table")
}

//验证用户信息
func checkFBusSubCM(u *FBusSubCM) (err error) {
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
	orm.RegisterModel(new(FBusSubCM))
}

//get node list
func GetFBusSubCMlist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FBusSubCM)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "FBusSubCM", "Description", "Cabinet", "Folder", "FileType", "Paramdesc", "ETLTaskDescId", "UpdateAt", "CreateAt")
	count, _ = qs.Count()
	return nodes, count
}

func GetFBusSubCMById(nid int64) (FBusSubCM, error) {
	o := orm.NewOrm()
	node := FBusSubCM{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddFBusSubCM(n *FBusSubCM) (int64, error) {
	if err := checkFBusSubCM(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(FBusSubCM)
	// node.DianPiaoCode = n.DianPiaoCode
	// node.DocOrg = n.DocOrg
	// node.BizdateId = n.BizdateId

	// node.DebitOrCredit = n.DebitOrCredit
	// node.ProductType = n.ProductType

	// node.Balance = n.Balance
	// node.Interest = n.Interest
	// node.InsidePlanFlag = n.InsidePlanFlag
	// node.StartAt = n.StartAt
	// node.OverAt = n.OverAt

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateFBusSubCM(n *FBusSubCM) (int64, error) {
	if err := checkFBusSubCM(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateFBusSubCM.FBusSubCM %q\n", n)
	o := orm.NewOrm()
	node := make(orm.Params)
	// if len(n.DianPiaoCode) > 0 {
	// node["DianPiaoCode"] = n.DianPiaoCode
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
	var table FBusSubCM
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelFBusSubCMById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&FBusSubCM{Id: Id})
	return status, err
}

func GetFBusSubCMlistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FBusSubCM)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetFBusSubCMTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(FBusSubCM)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllFBusSubCM retrieves all FBusSubCM matches certain condition. Returns empty list if
// no records exist
func GetAllFBusSubCMs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FBusSubCM))
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

	var l []FBusSubCM
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
