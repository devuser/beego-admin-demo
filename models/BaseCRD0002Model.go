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

//CRD0002Model
type CRD0002Model struct {
	Id int64 `orm:"column(id);pk;auto"`

	//组织机构编码
	Orgcode string `orm:"column(org_code);size(32);null"`

	// 业务形式
	Yewuxingshi string `orm:"column(yewuxingshi);size(255);"`
	// 贷款种类
	Daikuanzhonglei string `orm:"column(daikuanzhonglei);size(255);"`
	// 	借款人
	Jiekuanren string `orm:"column(jiekuanren);size(255);"`
	// 	合同编号
	Hetongbianhao string `orm:"column(hetongbianhao);size(32);"`
	// 	借据编号
	jiejubianhao string `orm:"column(jiejubianhao);size(32);"`
	// 	合同金额
	Hetongjine float64 `orm:"column(hetongjine);type(float64-decimal);decimals(2);digits(23);default(0)"`
	// 	发放金额
	Fafangjine float64 `orm:"column(fafangjine);type(float64-decimal);decimals(2);digits(23);default(0)"`
	// 	余额
	Yue float64 `orm:"column(yue);type(float64-decimal);decimals(2);digits(23);default(0)"`
	// 	合同利率（年%）
	Hetonglilv string `orm:"column(hetonglilv);size(32);"`
	// 现行利率（年%）

	// 	现行利率（年%）
	Xianxinglilv string `orm:"column(xianxinglilv);size(32);"`
	// 	利率浮动方式
	Lilvfudongfangshi string `orm:"column(lilvfudongfangshi);size(32);"`
	// 利率浮动值（%）
	Lilvfudongzhi string `orm:"column(lilvfudongzhi);size(32);"`
	// 	担保形式
	Danbaoxingshi string `orm:"column(danbaoxingshi);size(32);"`
	// 	担保单位或担保物
	Danbaodanweihuodanbaowu string `orm:"column(danbaodanweihuodanbaowu);size(32);"`
	// 	合同生效日
	Hetongshengxiaori string `orm:"column(hetongshengxiaori);size(8);"`
	// 	合同到期日
	Hetongdaoqiri string `orm:"column(hetongdaoqiri);size(8);"`
	// 	发放日
	Fafangri string `orm:"column(fafangri);size(8);"`
	// 	到期日
	Daoqiri string `orm:"column(daoqiri);size(8);"`
	// 	合同期限（月）
	Hetongqixian uint64 `orm:"column(hetongqixian);size(20)"`
	// 	借据期限（月）
	Jiejuqixian uint64 `orm:"column(jiejuqixian);size(20)"`
	// 	借款用途
	Jiekuanyongtu string `orm:"column(jiekuanyongtu);size(32);"`
	// 	五级分类结果
	Wujifenleijieguo string `orm:"column(wujifenleijieguo);size(32);"`
	// 	投向结构

	Touxiangjiegou string `orm:"column(touxiangjiegou);size(32);"`
	// 	是否逾期
	Shifouyuqi bool `orm:"column(shifouyuqi);size(1);default(0)"`

	// 	"企业分类
	Qiyefenlei string `orm:"column(qiyefenlei);size(32);"`
	////

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

func (n *CRD0002Model) TableName() string {
	return beego.AppConfig.String("base_crd_0002_table")
}

//验证用户信息
func checkCRD0002Model(u *CRD0002Model) (err error) {
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
	orm.RegisterModel(new(CRD0002Model))
}

//get node list
func GetCRD0002Modellist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(CRD0002Model)
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

func GetCRD0002ModelById(nid int64) (CRD0002Model, error) {
	o := orm.NewOrm()
	node := CRD0002Model{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddCRD0002Model(n *CRD0002Model) (int64, error) {
	if err := checkCRD0002Model(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(CRD0002Model)
	// node.Name = n.Name
	// node.State = n.State
	// node.Country = n.Country

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateCRD0002ModelById(n *CRD0002Model) (int64, error) {
	if err := checkCRD0002Model(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := make(orm.Params)
	//@todo 校验条件
	// if len(node) == 0 {
	// 	return 0, errors.New("update field is empty")
	// }
	var table CRD0002Model
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelCRD0002ModelById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&CRD0002Model{Id: Id})
	return status, err
}

func GetCRD0002ModellistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(CRD0002Model)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetCRD0002ModelTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(CRD0002Model)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllDownloadtrac retrieves all Downloadtrac matches certain condition. Returns empty list if
// no records exist
func GetAllCRD0002Model(query map[string]string, fields []string, sortby []string, order []string,
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
