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

// BaseKehu 客户
type BaseKehu struct {
	Id int64 `orm:"column(id);pk;auto"`

	//客户号
	KeHuHao string `orm:"column(kehuhao);unique;size(255)"`
	//客户中文名称
	Kehuzhwm string `orm:"column(kehuzhwm);size(255);null"`

	//归属机构
	GuishjigCode string `orm:"column(guishjig_code);size(32);null"`
	//服务机构
	ChjianjgCode string `orm:"column(chjianjgcode_code);size(32);null"`

	//所属集团编码
	HostGroupCode string `orm:"column(hostgroupcode);size(3);null"`
	//客户类型编码
	KehuLeiXing string `orm:"column(kehuleixing);size(255);null"`
	//客户分类编码
	KehuFenLei string `orm:"column(kehufenlei);size(255);null"`

	//企业客户法人性质
	QiYeKeHuFaRenXingZhi string `orm:"column(qiyekehufarenxingzhi);size(255);null"`
	// 企业分类（大中小微） E_QIYEGUIM
	Qiyeguim string `orm:"column(qiyeguim);size(255);null"`

	// 法人类别
	FaRenLeiBie string `orm:"column(farenleibie);size(255);null"`
	// 是否同业
	//	01 对公
	//	02 同业
	Kehuleix string `orm:"column(kehuleix);default(02);null"`
	// 同业客户法人性质
	TongYeKeHuFaRenXingZhi string `orm:"column(tongyekehufarenxingzhi);size(255);null"`
	// 客户标签	客户标签	varchar(255)	255		FALSE	FALSE	FALSE
	// 是否封存
	FengCunFlag bool `orm:"column(FengCun_flag);default(0);null"`

	// 创建日期
	ChuangJianRiQI string `orm:"column(chuangjianriqi);size(8)"`

	// 证件类型	证件类型	varchar(255)	255		FALSE	FALSE	F	ALSE
	ZhengJianLX string `orm:"column(zhengjianlx);size(255);null"`
	// 组织机构代码证
	ZuZhiJiGouDaiMaZ string `orm:"column(zuzhijigoudaimaz);size(255);null"`
	//客户并表标识
	KeHuBingBiaoFlag bool `orm:"column(kehubingbiaoflag);null"`

	// 常用客户标签1	常用客户标签1	<Undefined>			FALSE	FALSE	FALSE
	// 常用客户标签2	常用客户标签2	<Undefined>			FALSE	FALSE	FALSE
	// 常用客户标签3	常用客户标签3	<Undefined>			FALSE	FALSE	FALSE

	// 客户状态
	Kehuztai string `orm:"column(kehuztai);size(1);null"`

	//同业客户类型_使用场景待核实
	TongYeKehuLX string `orm:"column(TongYeKehuLX);size(4);null"`

	//证照种类编码
	ZhengZhaoZhongLei string `orm:"column(ZhengZhaoZhongLeiId);size(4);null"`
	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

const (
	cstKEHULEIX_TONGYE = "00"
)

func (t *BaseKehu) TableName() string {
	return beego.AppConfig.String("base_common_kehu_table")
}

func init() {
	orm.RegisterModel(new(BaseKehu))
}

// AddBasekehu insert a new BaseKehu into database and returns
// last inserted Id on success.
func AddBasekehu(m *BaseKehu) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetBaseKehuById retrieves BaseKehu by Id. Returns error if
// Id doesn't exist
func GetBaseKehuById(id int64) (v *BaseKehu, err error) {
	o := orm.NewOrm()
	v = &BaseKehu{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

//IsTongye 是否同业
func (t *BaseKehu) IsTongye() bool {
	return t.KehuLeiXing == cstKEHULEIX_TONGYE
}

func GetBaseKehuList(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseKehu)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}

	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "KeHuHao",
		"Kehuzhwm", "Guishjig", "Chjianjg", "HostGroupCode",

		"KehuLeiXing",

		"KehuFenLei",

		"QiYeKeHuFaRenXingZhi",

		"Qiyeguim",

		"FaRenLeiBie",
		"Kehuleix",

		"TongYeKeHuFaRenXingZhi",
		"FengCunFlag",
		"ChuangJianRiQI",
		"ZhengJianLX",
		"ZuZhiJiGouDaiMaZ",
		"KeHuBingBiaoFlag",
		"Kehuztai",
		"TongYeKehuLX",
		"ZhengZhaoZhongLei",
		"CreateAt",
		"UpdateAt",
		"Status")
	count, _ = qs.Count()
	return nodes, count
}

// GetAllBaseKehu retrieves all BaseKehu matches certain condition. Returns empty list if
// no records exist
func GetAllBaseKehus(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(BaseKehu))
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

	var l []BaseKehu
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

// UpdateBasekehu updates BaseKehu by Id and returns error if
// the record to be updated doesn't exist
func UpdateBasekehuById(m *BaseKehu) (err error) {
	o := orm.NewOrm()
	v := BaseKehu{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteBaseKehu deletes BaseKehu by Id and returns error if
// the record to be deleted doesn't exist
func DelBasekehuById(id int64) (err error) {
	o := orm.NewOrm()
	v := BaseKehu{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&BaseKehu{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
