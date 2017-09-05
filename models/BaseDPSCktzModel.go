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

// DPSCkzhtz 存款账户台账
type DPSCkzhtz struct {
	Id int64 `orm:"column(id);pk;auto"`
	//客户号
	KeHuHaoo string `orm:"column(kehuhaoo);size(32);null"`
	//客户中文名称
	Kehuzhwm string `orm:"column(kehuzhwm);size(255);null"`
	//归属机构Id
	Guishjig string `orm:"column(guishjig_code);size(32);null"`
	//服务机构Id
	Chjianjg string `orm:"column(chjianjg_code);size(32);null"`
	//所属集团编码
	//长亮: E_SUOSHJT2 所属集团2
	HostGroupCode string `orm:"column(hostgroup_code);size(3);null"`
	//客户类型编码 E_KEHULEIX
	KehuLeiXing string `orm:"column(kehuleixing);size(255);null"`
	//客户分类编码
	KehuFenLei string `orm:"column(kehufenlei);size(255);null"`
	//账户号
	ZhangHao string `orm:"column(zhanghao);size(64);null"`
	//账户名称
	Zhuzwmc string `orm:"column(zhhuzwmc);size(255);null"`
	//子账号（序号）
	Zhhaoxuh string `orm:"column(zhhaoxuh);size(32);null"`
	//开户机构Id
	KaiHuOrgcode string `orm:"column(kaihuorg_code);size(32);null"`

	//开户日期
	Kaihriqi string `orm:"unique;column(kaihriqi);size(8);"`
	//销户日期
	Xiohriqi string `orm:"unique;column(xiohriqi);size(8);"`

	//存款类型_七种
	CunKuanLeiXing string `orm:"column(cunkuanleixing);size(4)"`

	//存款性质
	CunKuanXingZhi string `orm:"column(cunkuanxingzhi);size(255);null"`
	//客户并表标识
	KeHuBingBiaoFlag bool `orm:"column(kehubingbiao_flag);default(0);null"`
	//账户并表标识
	ZhangHuBingBiaoFlag bool `orm:"column(zhanghubingbiao_flag);default(0);null"`
	//执行利率，精度参考长亮
	ZhiXingLilv float64 `orm:"column(zhixinglilv);type(float64-decimal);decimals(6);digits(23);null"`
	//存款产品Id
	CunKuanChanPinId int64 `orm:"column(cunkuanchanpin);size(20);null"`
	//期限Id
	QiXian int64 `orm:"column(qixian);size(20);null"`
	//期限单位Id
	QiXianUnitId int64 `orm:"column(qixianunit_id);size(20);default(3)"`

	//银行集团账户标识（是否影子账户）
	YingHangJTZhangHUFlag bool `orm:"column(yinghangjtzhanghu_flag);default(0);null"`
	//货币代号
	Huobdaih string `orm:"column(huobdaih);type(string);size(2);default(01)"`

	//============= 与账户标签有关开始 =============
	//基本保险标签
	JiBenBaoXianFlag bool `orm:"column(jibenbaoxian_flag);default(0);null"`
	//补充医保标签
	BuChongYiBaoFlag bool `orm:"column(buchongyibao_flag);default(0);null"`

	//年金标签
	NianJinFlag bool `orm:"column(nianjin_flag);default(0);null"`
	//补充公积金标签
	BuChongGongJiJinFlag bool `orm:"column(buchonggongjijin_flag);default(0);null"`
	//党团工会标签
	DangTuanGongHuiFlag bool `orm:"column(dangtuangonghui_flag);default(0);null"`
	//其他标签
	QiTaFlag bool `orm:"column(qita_flag);default(0);null"`

	//同业活期保证金存款标签 待核实报表是否需要
	TYHuoQiBaoZhengJinFlag bool `orm:"column(tyhuoqibaozhengjin_flag);default(0);null"`
	//同业定期保证金存款 待核实报表是否需要
	TYDingQiBaoZhengJinFlag bool `orm:"column(tydingqibaozhengjin_flag);default(0);null"`
	//对公活期保证金存款 待核实报表是否需要
	DGHuoQiBaoZhengJinFlag bool `orm:"column(dghuoqibaozhengjin_flag);default(0);null"`
	//对公定期保证金存款 待核实报表是否需要
	DGDingQiBaoZhengJinFlag bool `orm:"column(dgdingqibaozhengjin_flag);default(0);null"`
	// <restrictionType id="E_ZHHUZTAI" longname="账户状态" base="BaseType.U_LEIXIN01" tags="LT">
	// 	<enumeration id="ZC" value="A" longname="正常"/>
	// 	<enumeration id="XH" value="C" longname="销户"/>
	// 	<enumeration id="JX" value="D" longname="久悬户"/>
	// 	<enumeration id="YW" value="I" longname="转营业外收入"/>
	// 	<enumeration id="JE" value="F" longname="金额冻结"/>
	// 	<enumeration id="FB" value="E" longname="封闭冻结"/>
	// </restrictionType>
	//账户状态
	Zhhuztai string `orm:"column(zhhuztai);type(string);size(1);default(A)"  form:"Zhhuztai"  valid:"Required"`

	//============= 与标签有有关结束 =============
	// Status int `orm:"column(status);default(2)" form:"Status"`
	//-以下是存储控制标志-	-以下是存储控制标志-	<Undefined>			FALSE	FALSE	FALSE
	//生效日期	生效日期
	ShengXiaoDate string `orm:"unique;column(shengxiaodate);size(8);"`
	ShiXiaoDate   string `orm:"unique;column(shixiaodate);size(8);"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (t *DPSCkzhtz) TableName() string {
	return beego.AppConfig.String("base_dps_cktz_table")
}

func init() {
	orm.RegisterModel(new(DPSCkzhtz))
}

// AddDPSCkzhtz insert a new DPSCkzhtz into database and returns
// last inserted Id on success.
func AddDPSCkzhtz(m *DPSCkzhtz) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetDPSCkzhtzById retrieves DPSCkzhtz by Id. Returns error if
// Id doesn't exist
func GetDPSCkzhtzById(id int64) (v *DPSCkzhtz, err error) {
	o := orm.NewOrm()
	v = &DPSCkzhtz{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetDPSCkzhtzList(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(DPSCkzhtz)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	//qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "KeHuHaoo", "Kehuzhwm", "Guishjig", "Chjianjg", "Status")
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "KeHuHaoo", "Kehuzhwm", "Status")
	count, _ = qs.Count()
	return nodes, count
}

// GetAllDPSCkzhtz retrieves all DPSCkzhtz matches certain condition. Returns empty list if
// no records exist
func GetAllDPSCkzhtzs(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(DPSCkzhtz))
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

	var l []DPSCkzhtz
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

// UpdateDPSCkzhtz updates DPSCkzhtz by Id and returns error if
// the record to be updated doesn't exist
func UpdateDPSCkzhtzById(m *DPSCkzhtz) (err error) {
	o := orm.NewOrm()
	v := DPSCkzhtz{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteDPSCkzhtz deletes DPSCkzhtz by Id and returns error if
// the record to be deleted doesn't exist
func DelDPSCkzhtzById(id int64) (err error) {
	o := orm.NewOrm()
	v := DPSCkzhtz{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&DPSCkzhtz{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
