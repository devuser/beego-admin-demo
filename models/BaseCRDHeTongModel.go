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

// CRDHeTong 借贷合同
type CRDHeTong struct {
	Id int64 `orm:"column(id);pk;auto"`

	//业务种类
	YeWuZhongLei string `orm:"column(yewuzhonglei);size(255);null"`
	//业务形式
	YeWuXingShi string `orm:"column(yewuxingshi);size(255);null"`
	//贷款形式
	DaiKuanXingShi string `orm:"column(daikuanxingshi);size(255);null"`
	//合同编号
	HeTongCode string `orm:"unique;column(hetongcode);size(255);null"`

	//服务机构
	Orgcode string `orm:"column(org_code);size(32);null"`

	//借款单位
	KeHuHaoo string `orm:"column(kehuhaoo);size(32);null"`
	//货币代号
	Huobdaih string `orm:"column(huobdaih);type(string);size(2);default(01)"`

	//合同金额，长亮对照 ZONGJINE	总金额
	HeTongJinE float64 `orm:"column(hetongjine);type(float64-decimal);decimals(2);digits(23);default(0)"`

	//利率
	LiLv float64 `orm:"column(lilv);type(float64-decimal);decimals(7);digits(23);default(0)"`

	//现行利率
	XianXingLiLv float64 `orm:"column(xianxinglilv);type(float64-decimal);decimals(7);digits(23);default(0)"`

	//合同利率（年%）
	Hetongll float64 `orm:"column(hetongll);type(float64-decimal);decimals(6);digits(11);default(0)"`
	// NYUELILV	年/月利率标识	varchar2(1)	Y		D--日利率
	// NYUELILV	年/月利率标识	varchar2(1)	Y		M--月利率
	// NYUELILV	年/月利率标识	varchar2(1)	Y		Y--年利率
	Nyuelilv string `orm:"column(nyuelilv);type(string);size(1);default(Y)"  form:"Nyuelilv"  valid:"Required"`
	//利率浮动方式
	Ratefloattype string `orm:"column(ratefloattype);size(2);null"`

	// 主要担保形式编码（五种）
	Vouchtype string `orm:"column(vouchtype);size(4);null"`
	//担保单位或担保物
	DanBaoDanWeiHuoDanBaoWu string `orm:"column(danbaodanweihuodanbaowu);size(255);null"`
	//合同签订日_使用场景待核实	合同签订日_使用场景待核实	date			FALSE	FALSE	FALSE
	HeTongQianDingRi time.Time `orm:"column(hetongqiandingri);null"`
	//合同生效日	合同生效日	date			FALSE	FALSE	FALSE
	HeTongShengXiaoRi string `orm:"unique;column(hetongshengxiaori);size(8);"`
	//合同到期日	合同到期日	date			FALSE	FALSE	FALSE
	HeTongDaoQiRi string `orm:"unique;column(hetongdaoqiri);size(8);"`
	//合同期限
	HeTongQiXian int64 `orm:"column(hetongqixian);size(20);null"`
	//合同期限单位
	HeTongQiXianUnit int64 `orm:"column(hetongqixianunit);size(20);null"`
	//借款用途
	JieKuanYongTu string `orm:"column(jiekuanyongtu);size(255);null"`
	// 五级分类结果                  (填空)
	WuJiFenLei string `orm:"column(wujifenlei);size(255);null"`
	// 投向结构
	TouXiangJieGou string `orm:"column(touxiangjiegou);size(255);null"`
	// 状态
	HeTongStatus string `orm:"column(hetongstatus);size(255);null"`

	//BAILSUM	保证金金额	NUMBER(32,4)
	Bailsum float64 `orm:"column(bailsum);type(float64-decimal);decimals(4);digits(32);default(0)"`
	//BAILRATIO	保证金比例	NUMBER(32,4)
	Bailratio float64 `orm:"column(bailratio);type(float64-decimal);decimals(4);digits(32);default(0)"`
	//PDGRATIO	手续费比例	NUMBER(32,4)
	Pdgratio float64 `orm:"column(pdgratio);type(float64-decimal);decimals(4);digits(32);default(0)"`
	//PDGSUM	手续费	NUMBER(32,4)
	Pdgsum float64 `orm:"column(pdgsum);type(float64-decimal);decimals(4);digits(32);default(0)"`
	//SETTLEBANK	清算行名称 ?	VARCHAR2(80)
	Settlebank string `orm:"column(settlebank);type(string);size(80);null"`
	//LEADBANK	牵头行	VARCHAR2(80)
	Leadbank string `orm:"column(LEADBANK);type(string);size(80);null"`

	//MNGFEERATE	管理费率?	NUMBER(32,4)
	Mngfeerate float64 `orm:"column(mngfeerate);type(float64-decimal);decimals(4);digits(32);default(0)"`

	//MNGFEE	管理费金额	NUMBER(32,4)
	Mngfee float64 `orm:"column(mngfee);type(float64-decimal);decimals(4);digits(32);default(0)"`

	// 手续费率
	ShouXuFeiLv float64 `orm:"column(shouxufeilv);type(float64-decimal);decimals(7);digits(23);default(0)"`
	// 是否委托
	WeiTuoFlag bool `orm:"column(weituo_flag);default(0);null"`
	// 委托人（和委托贷款有关）
	WeiTuoRen string `orm:"column(weituoren);size(255);null"`
	// 是否循环
	Xunhdaik bool `orm:"column(xunhdaik);default(0)"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *CRDHeTong) TableName() string {
	return beego.AppConfig.String("base_crd_jiedaihetong_table")
}

func init() {
	orm.RegisterModel(new(CRDHeTong))
}

// AddCRDHeTong insert a new CRDHeTong into database and returns
// last inserted Id on success.
func AddCRDHeTong(m *CRDHeTong) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// Id doesn't exist
func GetCRDHeTongById(id int64) (v *CRDHeTong, err error) {
	o := orm.NewOrm()
	v = &CRDHeTong{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCRDHeTong retrieves all CRDHeTong matches certain condition. Returns empty list if
// no records exist
func GetAllCRDHeTong(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CRDHeTong))
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

	var l []CRDHeTong
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

// UpdateCRDHeTong updates CRDHeTong by Id and returns error if
// the record to be updated doesn't exist
func UpdateCRDHeTongById(m *CRDHeTong) (err error) {
	o := orm.NewOrm()
	v := CRDHeTong{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCRDHeTong deletes CRDHeTong by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCRDHeTong(id int64) (err error) {
	o := orm.NewOrm()
	v := CRDHeTong{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CRDHeTong{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
