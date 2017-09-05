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

// FefpLoadM 融资表
type FefpLoadM struct {
	Id                        int64     `orm:"column(id);auto;pk"`
	Loancode                  string    `orm:"size(20);"`
	Loanccy                   float64   `orm:"type(float64-decimal);default(0)"`
	Loanamount                float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Loanterm                  float64   `orm:"type(float64-decimal);default(0)"`
	Loangrace                 float64   `orm:"type(float64-decimal);default(0)"`
	Loandate                  time.Time `orm:"type(datetime);"`
	Loanduedate               time.Time `orm:"type(datetime);"`
	Loanaccount               string    `orm:"size(30);"`
	Settingaccount            string    `orm:"size(30);"`
	Loanflag                  string    `orm:"size(255);"`
	Accountccy                float64   `orm:"type(float64-decimal);default(0)"`
	Exchangerate              float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Subcontractid             float64   `orm:"type(float64-decimal);default(0)"`
	Incomeway                 string    `orm:"size(255);"`
	Settingway                string    `orm:"size(255);"`
	Ratetype                  string    `orm:"size(255);"`
	Fixrate                   float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Rateid                    float64   `orm:"type(float64-decimal);default(0)"`
	Floatcycleunit            string    `orm:"size(255);"`
	Floatcycle                float64   `orm:"type(float64-decimal);default(0)"`
	Addsubway                 string    `orm:"size(255);"`
	Addsubvalue               float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Lockingflag               string    `orm:"size(255);"`
	Batchid                   float64   `orm:"type(float64-decimal);default(0)"`
	Transtype                 float64   `orm:"type(float64-decimal);default(0)"`
	Transitemid               float64   `orm:"type(float64-decimal);default(0)"`
	Extendflag                string    `orm:"size(255);"`
	Extendterm                float64   `orm:"type(float64-decimal);default(0)"`
	Loanbalance               float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Receivableinterest        float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Receivedinterest          float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Accruedinterest           float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Currentinterest           float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Lastinterestset           float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Currentinterestset        float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Currentamortizedinterest  float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Currentaccruedinterest    float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Overdueamount             float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Overduepayamount          float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Overdueinterest           float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Currenttransoverdueamount float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Currentoverdueinterest    float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Lessbalance               float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Currentlessamount         float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Lastbearingondate         float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	State                     float64   `orm:"type(float64-decimal);default(0)"`
	Opstate                   float64   `orm:"type(float64-decimal);default(0)"`
	WorkItemId                float64   `orm:"type(float64-decimal);default(0)"`
	Addsubtype                string    `orm:"size(255);"`
	Endflag                   string    `orm:"size(255);"`
	Txno                      float64   `orm:"type(float64-decimal);default(0)"`
	T_payamount               float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	T_payreceivable           float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	T_payaccrued              float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	T_paycurrent              float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Loanscale                 float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Lastfixdate               time.Time `orm:"type(datetime);"`
	Sellerrefcode             string    `orm:"size(60);"`
	Loantype                  string    `orm:"size(255);"`
	Lastmoduser               string    `orm:"size(50);"`
	Lastmoddate               time.Time `orm:"type(datetime);"`
	Paiditem                  string    `orm:"size(255);"`
	Zqfee_rate                float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	T_pay_total               float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`
	Jjmflag                   string    `orm:"size(255);"`
	Expiredflag               string    `orm:"size(255);"`
	Expiredrate               float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	Interratetimezoneflag     float64   `orm:"type(float64-decimal);default(0)"`
	T_temppooluse             float64   `orm:"type(float64-decimal);decimals(2);digits(23);default(0)"`

	//业务日期
	bizdate string `orm:"column(bizdate);size(8)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (n *FefpLoadM) TableName() string {
	return beego.AppConfig.String("base_fefp_loanm_table")
}

//验证用户信息
func checkFefpLoadM(u *FefpLoadM) (err error) {
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
	orm.RegisterModel(new(FefpLoadM))
}

//get node list
func GetFefpLoadMlist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FefpLoadM)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "FefpLoadM", "Description", "Cabinet", "Folder", "FileType", "Paramdesc", "ETLTaskDescId", "UpdateAt", "CreateAt")
	count, _ = qs.Count()
	return nodes, count
}

func GetFefpLoadMById(nid int64) (FefpLoadM, error) {
	o := orm.NewOrm()
	node := FefpLoadM{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddFefpLoadM(n *FefpLoadM) (int64, error) {
	if err := checkFefpLoadM(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(FefpLoadM)
	// node.DianPiaoCode = n.DianPiaoCode
	// node.DocOrg = n.DocOrg
	// node.BizdateId = n.BizdateId
	//
	// node.DebitOrCredit = n.DebitOrCredit
	// node.ProductType = n.ProductType
	//
	// node.Balance = n.Balance
	// node.Interest = n.Interest
	// node.InsidePlanFlag = n.InsidePlanFlag
	// node.StartAt = n.StartAt
	// node.OverAt = n.OverAt

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateFefpLoadM(n *FefpLoadM) (int64, error) {
	if err := checkFefpLoadM(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateFefpLoadM.FefpLoadM %q\n", n)
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
	var table FefpLoadM
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelFefpLoadMById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&FefpLoadM{Id: Id})
	return status, err
}

func GetFefpLoadMlistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FefpLoadM)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetFefpLoadMTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(FefpLoadM)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllFefpLoadM retrieves all FefpLoadM matches certain condition. Returns empty list if
// no records exist
func GetAllFefpLoadMs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FefpLoadM))
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

	var l []FefpLoadM
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
