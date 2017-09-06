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

// private String buBen;
// //  	上月末余额
// private BigDecimal shangYueMoYue;
//
// // 	上日余额
// private BigDecimal shangRiYue;
//
// // 	核定贷款
// private BigDecimal heDingDaiKuan;
// // 	截至上日月均
// private BigDecimal jieZhiShangRiYueJun;
//
// // 	截至上日累计日均
// private BigDecimal jieZhiShangRiLeiJiRiJun;
//
// //		当日发生		当日余额
//
// //    发放
// private BigDecimal dangRiFaShengFaFang;
//
// // 回收
// private BigDecimal dangRiFaShengHuiShou;
//
// // 	实际
// private BigDecimal dangRiFaShengShiJi;
// //    当日余额
// private BigDecimal dangRiYue;
//CRD0001Model 担保形式
type CRD0001Model struct {
	Id int64 `orm:"column(id);pk;auto"`

	//组织机构编码
	Orgcode string `orm:"column(org_code);size(32);null"`
	//上月末余额
	Shangyuemoyue float64 `orm:"column(shangyuemoyue);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//上日余额
	Shangriyue float64 `orm:"column(shangriyue);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//
	//核定贷款
	Hedingdaikuan float64 `orm:"column(hedingdaikuan);type(float64-decimal);decimals(2);digits(23);default(0)"`
	// 截至上日月均
	Jiezhishangriyuejun float64 `orm:"column(jiezhishangriyuejun);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//截至上日累计日均
	Jiezhishangrileijirijun float64 `orm:"column(jiezhishangrileijirijun);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//当日发生		当日余额
	//发放
	Dangrifashengfafang float64 `orm:"column(dangrifashengfafang);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//
	//回收
	Dangrifashenghuishou float64 `orm:"column(dangrifashenghuishou);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//实际
	Dangrifashengshiji float64 `orm:"column(dangrifashengshiji);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//当日余额
	Dangriyue float64 `orm:"column(dangriyue);type(float64-decimal);decimals(2);digits(23);default(0)"`

	Bizdate string `orm:"column(bizdate);size(8);"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

// TableName 缺省方法，返回对应的数据库表名
func (n *CRD0001Model) TableName() string {
	return beego.AppConfig.String("base_crd_0001_table")
}

//验证用户信息
func checkCRD0001Model(u *CRD0001Model) (err error) {
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
	orm.RegisterModel(new(CRD0001Model))
}

//get node list
func GetCRD0001Modellist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(CRD0001Model)
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

func GetCRD0001ModelById(nid int64) (CRD0001Model, error) {
	o := orm.NewOrm()
	node := CRD0001Model{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddCRD0001Model(n *CRD0001Model) (int64, error) {
	if err := checkCRD0001Model(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(CRD0001Model)
	// node.Name = n.Name
	// node.State = n.State
	// node.Country = n.Country

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateCRD0001ModelById(n *CRD0001Model) (int64, error) {
	if err := checkCRD0001Model(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := make(orm.Params)
	//@todo 校验条件
	// if len(node) == 0 {
	// 	return 0, errors.New("update field is empty")
	// }
	var table CRD0001Model
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelCRD0001ModelById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&CRD0001Model{Id: Id})
	return status, err
}

func GetCRD0001ModellistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(CRD0001Model)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetCRD0001ModelTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(CRD0001Model)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllDownloadtrac retrieves all Downloadtrac matches certain condition. Returns empty list if
// no records exist
func GetAllCRD0001Model(query map[string]string, fields []string, sortby []string, order []string,
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
