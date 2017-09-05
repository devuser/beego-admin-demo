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

// FBusCTCTM 国内保理-合同表
type FBusCTCTM struct {
	Id int64 `orm:"column(id);auto;pk"`

	SellerId string `orm:"column(seller_id);size(32);unique" form:"SellerId" valid:"Required"`
	//@TODO 待核实
	ProductId  float64 `orm:"column(product_id);type(float64-decimal);default(0)"`
	Contractno string  `orm:"column(contractno);size(30);"`

	Startdate time.Time `orm:"column(start_date);type(datetime);"`
	Enddate   time.Time `orm:"column(end_date);type(datetime);"`

	Contractccy string  `orm:"column(contractccy);size(32);"`
	Contractamt float64 `orm:"column(contractamt);type(float64-decimal);decimals(8);digits(39);default(0)"`

	Inactflag       string    `orm:"column(inactflag);size(1);"`
	Inactdate       time.Time `orm:"column(inactdate);type(datetime);"`
	CycleFlag       string    `orm:"column(cycleflag);size(1);"`
	ShareFlag       string    `orm:"column(shareflag);size(1);"`
	RebchId         string    `orm:"column(rebchid);size(32);"`
	Custmoermgr     string    `orm:"column(custmoermgr);size(30);"`
	InterincomeFlag string    `orm:"column(interincomeflag);size(1);"`
	IntercloseFlag  string    `orm:"column(intercloseflag);size(1);"`
	InterpayFlag    string    `orm:"column(interpayflag);size(1);"`
	RateFlag        string    `orm:"column(rateflag);size(10);"`
	InterrateFlag   string    `orm:"column(interrateflag);size(1);"`
	//@TODO 待核实
	InterratetimezoneFlag float64 `orm:"column(interratetimezoneflag);type(float64-decimal);decimals(8);digits(39);default(0)"`
	DayormonthFlag        string  `orm:"column(dayormonthflag);size(1);"`
	DayormonthValue       float64 `orm:"column(dayormonthvalue);type(float64-decimal);default(0)"`
	InterestRate          float64 `orm:"column(interestrate);type(float64-decimal);decimals(8);digits(39);default(0)"`
	LockFlag              string  `orm:"column(lockflag);size(1);"`
	JjmFlag               string  `orm:"column(jjmflag);size(1);"`
	JjmmodFlag            string  `orm:"column(jjmmodflag);size(1);"`
	JjmRate               float64 `orm:"column(jjmrate);type(float64-decimal);decimals(8);digits(39);default(0)"`
	ExpiredFlag           string  `orm:"column(expiredflag);size(1);"`
	ExpiredRate           float64 `orm:"column(expiredrate);type(float64-decimal);decimals(8);digits(39);default(0)"`
	FeeaccFlag            string  `orm:"column(feeaccflag);size(1);"`
	FeepaidFlag           string  `orm:"column(feepaidflag);size(1);"`
	FeeclFlag             string  `orm:"column(feeclflag);size(1);"`
	FeeamtFlag            string  `orm:"column(feeamtflag);size(1);"`
	//@TODO 待核实
	OpState float64 `orm:"column(opstate);type(float64-decimal);default(0)"`
	//@TODO 待核实
	WorkItemId float64 `orm:"column(workitemid);type(float64-decimal);default(0)"`

	State float64 `orm:"column(state);type(float64-decimal);default(0)"`
	Zyqx  float64 `orm:"column(zyqx);type(float64-decimal);default(0)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (n *FBusCTCTM) TableName() string {
	return beego.AppConfig.String("base_fbus_ctctm_table")
}

//验证用户信息
func checkFBusCTCTM(u *FBusCTCTM) (err error) {
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
	orm.RegisterModel(new(FBusCTCTM))
}

//get node list
func GetFBusCTCTMlist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FBusCTCTM)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "FBusCTCTM", "Description", "Cabinet", "Folder", "FileType", "Paramdesc", "ETLTaskDescId", "UpdateAt", "CreateAt")
	count, _ = qs.Count()
	return nodes, count
}

func GetFBusCTCTMById(nid int64) (FBusCTCTM, error) {
	o := orm.NewOrm()
	node := FBusCTCTM{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddFBusCTCTM(n *FBusCTCTM) (int64, error) {
	if err := checkFBusCTCTM(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(FBusCTCTM)
	// node.DianPiaoCode = n.DianPiaoCode
	// node.DocOrg = n.DocOrg
	// node.BizdateId = n.BizdateId
	//
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
func UpdateFBusCTCTM(n *FBusCTCTM) (int64, error) {
	if err := checkFBusCTCTM(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateFBusCTCTM.FBusCTCTM %q\n", n)
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
	var table FBusCTCTM
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelFBusCTCTMById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&FBusCTCTM{Id: Id})
	return status, err
}

func GetFBusCTCTMlistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FBusCTCTM)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetFBusCTCTMTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(FBusCTCTM)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllFBusCTCTM retrieves all FBusCTCTM matches certain condition. Returns empty list if
// no records exist
func GetAllFBusCTCTMs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FBusCTCTM))
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

	var l []FBusCTCTM
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
