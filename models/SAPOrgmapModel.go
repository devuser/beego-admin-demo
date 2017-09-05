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

//SAPOrgmap SAP组织机构对照表
//
//	末端机构存在一一对应关系
//	非末端机构，没有SAP机构编码
//	非末端机构建议方案1 根据doc_orgs表中维护的组织机构树关系，合并下属的末端机构科目余额
//	非末端机构建议方案2 根据本表中维护的组织机构树关系，合并下属的末端机构科目余额
//	目前执行建议方案2
type SAPOrgmap struct {
	Id int64 `orm:"column(id);pk;auto"`
	//核心组织机构编码
	Orgcode string `orm:"column(org_code);unique;size(32);valid(required)"`
	//是否末端机构
	LeafOrg bool `orm:"column(leaf_org);default(1)"`
	//SAP机构编码
	SAPOrgcode string `orm:"column(sap_org_code);size(32);default()"`
	//上级机构编码
	ParentOrgcode string `orm:"column(parent_org_code);size(32);"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *SAPOrgmap) TableName() string {
	return beego.AppConfig.String("base_saporgmap_table")
}

func init() {
	orm.RegisterModel(new(SAPOrgmap))
}

// AddSAPOrgmap insert a new SAPOrgmap into database and returns
// last inserted Id on success.
func AddSAPOrgmap(m *SAPOrgmap) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetSAPOrgmapById retrieves SAPOrgmap by Id. Returns error if
// Id doesn't exist
func GetSAPOrgmapById(id int64) (v *SAPOrgmap, err error) {
	o := orm.NewOrm()
	v = &SAPOrgmap{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//get node list
func GetSAPOrgmaplist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(SAPOrgmap)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}

	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id",
		"Orgcode",
		"LeafOrg",
		"SAPOrgcode",
		"ParentOrgcode",
	)
	count, _ = qs.Count()
	return nodes, count
}

// GetAllSAPOrgmap retrieves all SAPOrgmap matches certain condition. Returns empty list if
// no records exist
func GetAllSAPOrgmap(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(SAPOrgmap))
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

	var l []SAPOrgmap
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
func checkSAPOrgmap(u *SAPOrgmap) (err error) {
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

func UpdateSAPOrgmap(n *SAPOrgmap) (int64, error) {
	if err := checkSAPOrgmap(n); err != nil {
		return 0, err
	}
	fmt.Printf("UpdateSAPOrgmap.Orgcode %s\n", n.Orgcode)
	fmt.Printf("UpdateSAPOrgmap.LeafOrg %s\n", n.LeafOrg)
	fmt.Printf("UpdateSAPOrgmap.SAPOrgcode %s\n", n.SAPOrgcode)

	o := orm.NewOrm()
	node := make(orm.Params)
	if len(n.Orgcode) > 0 {
		node["Orgcode"] = n.Orgcode

	}
	node["SAPOrgcode"] = n.SAPOrgcode
	node["LeafOrg"] = n.LeafOrg
	if !n.LeafOrg {
		node["SAPOrgcode"] = ""
	}
	node["ParentOrgcode"] = n.ParentOrgcode
	if len(node) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table SAPOrgmap
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

// DeleteSAPOrgmap deletes SAPOrgmap by Id and returns error if
// the record to be deleted doesn't exist

func DelSAPOrgmapById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&SAPOrgmap{Id: Id})
	return status, err
}
