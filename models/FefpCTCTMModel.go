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

// FefpCTCTM 池融资
type FefpCTCTM struct {
	Id                    int64     `orm:"column(id);auto;pk"`
	SellerId              string    `orm:"size(60);"`
	ProductId             float64   `orm:"type(float64-decimal);default(0)"`
	Contractno            string    `orm:"size(30);"`
	Startdate             string    `orm:"column(startdate);size(8)"`
	Enddate               string    `orm:"column(enddate);size(8)"`
	Contractccy           string    `orm:"size(2);"`
	Contractamt           float64   `orm:"type(float64-decimal);decimals(8);digits(39);default(0)"`
	InactFlag             string    `orm:"size(255);"`
	Inactdate             string    `orm:"column(inactdate);size(8)"`
	CycleFlag             string    `orm:"size(255);"`
	ShareFlag             string    `orm:"size(255);"`
	RebchId               float64   `orm:"type(float64-decimal);default(0)"`
	Custmoermgr           string    `orm:"size(30);"`
	InterincomeFlag       string    `orm:"size(255);"`
	IntercloseFlag        string    `orm:"size(255);"`
	InterpayFlag          string    `orm:"size(255);"`
	InterrateFlag         string    `orm:"size(255);"`
	InterratetimezoneFlag float64   `orm:"type(float64-decimal);default(0)"`
	DayormonthFlag        string    `orm:"size(255);"`
	Dayormonthvalue       float64   `orm:"type(float64-decimal);default(0)"`
	Interestrate          float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	LockFlag              string    `orm:"size(255);"`
	JjmFlag               string    `orm:"size(255);"`
	JjmmodFlag            string    `orm:"size(255);"`
	Jjmrate               float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	ExpiredFlag           string    `orm:"size(255);"`
	Expiredrate           float64   `orm:"type(float64-decimal);decimals(4);digits(23);default(0)"`
	FeeaccFlag            string    `orm:"size(255);"`
	FeepaidFlag           string    `orm:"size(255);"`
	FeeclFlag             string    `orm:"size(255);"`
	FeeamtFlag            string    `orm:"size(255);"`
	Opstate               float64   `orm:"type(float64-decimal);default(0)"`
	WorkItemId            float64   `orm:"type(float64-decimal);default(0)"`
	State                 float64   `orm:"type(float64-decimal);default(0)"`
	Zyqx                  float64   `orm:"type(float64-decimal);default(0)"`
	FpoolproductId        float64   `orm:"type(float64-decimal);default(0)"`
	Lastmoduser           string    `orm:"size(50);"`
	Lastmoddate           string    `orm:"column(lastmoddate);size(8)"`
	Txno                  string    `orm:"size(50);"`
	Creatdate             time.Time `orm:"type(datetime);null"`
	AccountId             string    `orm:"size(50);"`
	Sellername            string    `orm:"size(60);"`
	Fpoolaccountamt       float64   `orm:"type(float64-decimal);decimals(8);digits(39);default(0)"`

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

func (n *FefpCTCTM) TableName() string {
	return beego.AppConfig.String("base_fefp_ctctm_table")
}

//验证用户信息
func checkFefpCTCTM(u *FefpCTCTM) (err error) {
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
	orm.RegisterModel(new(FefpCTCTM))
}

//get node list
func GetFefpCTCTMlist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FefpCTCTM)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "FefpCTCTM", "Description", "Cabinet", "Folder", "FileType", "Paramdesc", "ETLTaskDescId", "UpdateAt", "CreateAt")
	count, _ = qs.Count()
	return nodes, count
}

func GetFefpCTCTMById(nId int64) (FefpCTCTM, error) {
	o := orm.NewOrm()
	node := FefpCTCTM{Id: nId}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddFefpCTCTM(n *FefpCTCTM) (int64, error) {
	if err := checkFefpCTCTM(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(FefpCTCTM)
	// node.DianPiaoCode = n.DianPiaoCode
	// node.DocOrg = n.DocOrg
	// node.BizdateId= n.BizdateId	//
	// node.DebitOrCredit = n.DebitOrCredit
	// node.ProductType = n.ProductType
	//
	// node.Balance = n.Balance
	// node.Interest = n.Interest
	// node.InsidePlanFlag = n.InsidePlanFlag 	// node.StartAt = n.StartAt
	// node.OverAt = n.OverAt

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateFefpCTCTM(n *FefpCTCTM) (int64, error) {
	if err := checkFefpCTCTM(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateFefpCTCTM.FefpCTCTM %q\n", n)
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
	var table FefpCTCTM
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelFefpCTCTMById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&FefpCTCTM{Id: Id})
	return status, err
}

func GetFefpCTCTMlistByGroupid(GroupId int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(FefpCTCTM)
	count, _ = o.QueryTable(node).Filter("Group", GroupId).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetFefpCTCTMTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(FefpCTCTM)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllFefpCTCTM retrieves all FefpCTCTM matches certain condition. Returns empty list if
// no records exist
func GetAllFefpCTCTMs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(FefpCTCTM))
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
					return nil, errors.New("Error: InvalIdorder. Must be either [asc|desc]")
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
					return nil, errors.New("Error: InvalIdorder. Must be either [asc|desc]")
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

	var l []FefpCTCTM
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
