package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"reflect"
	"strings"
	"time"
)

// CRDJieJu 借贷合同
type CRDJieJu struct {
	Id int64 `orm:"column(id);pk;auto"`
	//合同编号
	HeTongCode string `orm:"column(hetongcode);size(255);null"`
	//借据编号
	JieJuCode string `orm:"unique;column(jiejucode);size(255);null"`

	//发放日
	FaFangRi string `orm:"unique;column(fafangri);size(8);"`
	// 到期日	到期日	date			FALSE	FALSE	FALSE
	DaoQiRi string `orm:"unique;column(daoqiri);size(8);"`
	//是否逾期	是否逾期	boolean			FALSE	FALSE	FALSE
	YuQiFlag bool `orm:"column(yuqi);null"`
	//发放金额（明确属于借据）	发放金额（明确属于借据）
	FangFangJE float64 `orm:"column(fangfangje);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//开户金融机构（使用场景待核实）	开户金融机构（使用场景待核实）
	KaiHuJinRongJiGou string `orm:"column(kaihujinrongjigou);size(255);null"`
	//客户在开户金融机构的账号（使用场景待核实）
	KaiHuJinRongJiGouKaiHuZhangHao string `orm:"column(kaihujinrongjigoukaihuzhanghao);size(255);null"`
	//借款金额	借款金额	decimal(19,2)	19	2	FALSE	FALSE	FALSE
	JieKuanJE float64 `orm:"column(jiekuanje);type(float64-decimal);decimals(2);digits(23);default(0)"`
	//付款方式
	FuKuanFangShi string `orm:"column(fukuanfangshi);size(255);null"`
	// 借款用途	借款用途	varchar(255)	255		FALSE	FALSE	FALSE
	JieKuanYongTu string `orm:"column(jiekuanyongtu);size(255);null"`
	//利率
	LiLv float64 `orm:"column(lilv);type(float64-decimal);decimals(7);digits(23);default(0)"`
	//提款日期
	TiKuanRQ string `orm:"unique;column(tikuanrq);size(8);"`
	//还款日期
	HuanKuanRQ string `orm:"unique;column(huankuanrq);size(8);"`
	//借据期限
	JieJuQiXian int64 `orm:"column(jiejuqixian);size(20);null"`
	//借据期限
	JieJuQiXianUnit int64 `orm:"column(jiejuqixianunit_id);size(20);default(3)"`
	//货币代号
	Huobdaih string `orm:"column(huobdaih);type(string);size(2);default(01)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

// TableName 缺省方法，返回对应的数据库表名
func (t *CRDJieJu) TableName() string {
	return beego.AppConfig.String("base_crd_jieju_table")
}

func init() {
	orm.RegisterModel(new(CRDJieJu))
}

// AddCRDJieJu insert a new CRDJieJu into database and returns
// last inserted Id on success.
func AddCRDJieJu(m *CRDJieJu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// Id doesn't exist
func GetCRDJieJuById(id int64) (v *CRDJieJu, err error) {
	o := orm.NewOrm()
	v = &CRDJieJu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRDJieJu retrieves all CRDJieJu matches certain condition. Returns empty list if
// no records exist
func GetAllCRDJieJu(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRDJieJu))
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

	var l []CRDJieJu
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

// UpdateCRDJieJu updates CRDJieJu by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRDJieJuById(m *CRDJieJu) (err error) {
	o := orm.NewOrm()
	v := CRDJieJu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRDJieJu deletes CRDJieJu by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRDJieJu(id int64) (err error) {
	o := orm.NewOrm()
	v := CRDJieJu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRDJieJu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
