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

//BaseSTMQudaoTranflow 渠道-企业转账指令表
//
//	企业转账指令表	CB_TRANFLOW	TFL
type BaseSTMQudaoTranflow struct {
	Id int64 `orm:"column(id);pk;auto"`
	//批次号
	Batchno string `orm:"column(batchno);size(15);null"`
	//网银流水号
	Flowno string `orm:"column(flowno);size(20);null"`
	//核心客户号 对应宇信内部字段 cif_hostno
	HostKehuhao string `orm:"column(host_kehuhao);size(20);null"`
	//网银客户号
	Cstno string `orm:"column(cstno);size(20);null"`
	//操作员编号
	Oprno string `orm:"column(oprno);size(20);null"`
	//操作员姓名
	Oprname string `orm:"column(oprname);size(60);null"`
	//交易代码
	Bsncode string `orm:"column(bsncode);size(8);null"`

	// 7	type	char	1	yes
	Type string `orm:"column(type);size(8);null"`
	// 8	payacc	varchar2	32	yes	付款账号
	Payacc string `orm:"column(payacc);size(32);null"`
	// 9	payname	varchar2	300	yes	付款人名称
	Payname string `orm:"column(payname);size(300);null"`
	// 10	paynode	varchar2	10	yes	付款人开户网点
	Paynode string `orm:"column(paynode);size(10);null"`
	// 11	rcvacc	varchar2	32	yes	收款账号
	Rcvacc string `orm:"column(rcvacc);size(32);null"`
	// 12	rcvname	varchar2	300	yes	收款人名称
	Rcvname string `orm:"column(rcvname);size(300);null"`
	// 13	rcvbank	varchar2	300	yes	收款行名称
	Rcvbank string `orm:"column(rcvbank);size(300);null"`
	// 14	comitrno	varchar2	12	yes
	Comitrno string `orm:"column(comitrno);size(12);null"`
	// 15	cstinnerno	varchar2	20	yes
	Cstinnerno string `orm:"column(cstinnerno);size(20);null"`
	// 16	totype	char	1	yes
	Totype string `orm:"column(totype);size(1);null"`
	// 17	tranchannel	varchar2	4	yes	交易渠道
	Tranchannel string `orm:"column(tranchannel);size(4);null"`
	// 18	cry	varchar2	3	yes	币种
	Cry string `orm:"column(cry);size(3);null"`
	// 19	tranamt	number	18,2	yes	交易金额
	Tranamt float64 `orm:"column(tranamt);type(float64-decimal);decimals(2);digits(23);default(0)"`
	// 20	purpose	varchar2	60	yes
	Purpose string `orm:"column(purpose);size(60);null"`
	// 21	rem	varchar2	300	yes
	Rem string `orm:"column(rem);size(300);null"`
	// 22	bankrem	varchar2	300	yes
	Bankrem string `orm:"column(bankrem);size(300);null"`
	// 23	submittime	char	14	no	提交时间
	Submittime string `orm:"column(submittime);size(14);null"`
	// 24	dealtime	char	14	yes	处理时间
	Dealtime string `orm:"column(dealtime);size(14);null"`
	// 25	fee1	number	18,2	yes
	Fee1 float64 `orm:"column(fee1);type(float64-decimal);decimals(2);digits(23);default(0)"`
	// 26	fee2	number	18,2	yes
	Fee2 float64 `orm:"column(fee2);type(float64-decimal);decimals(2);digits(23);default(0)"`
	// 27	rmttype	char	1	yes
	Rmttype string `orm:"column(rmttype);size(1);null"`
	// 28	chkcode	varchar2	30	yes
	Cshkcode string `orm:"column(chkcode);size(30);null"`
	// 29	hostcode	varchar2	20	yes
	Hostcode string `orm:"column(hostcode);size(20);null"`
	// 30	sendtime	char	14	yes
	Sendtime string `orm:"column(sendtime);size(14);null"`
	// 31	stt	char	2	no
	Stt string `orm:"column(stt);size(2);null"`
	// 32	appid	varchar2	20	yes
	AppId string `orm:"column(appid);size(20);null"`
	// 33	hostflowno	varchar2	32	yes
	HostFlowno string `orm:"column(hostflowno);size(32);null"`
	// 34	sendhostno	varchar2	32	yes
	SendHostno string `orm:"column(sendhostno);size(32);null"`
	// 35	errmsg	varchar2	300	yes
	Errmsg string `orm:"column(errmsg);size(300);null"`
	// 36	refusereason	varchar2	300	yes	拒绝原因
	RefuseReason string `orm:"column(refusereason);size(300);null"`
	// 37	orderdealtime	char	14	yes
	OrderDealTime string `orm:"column(orderdealtime);size(14);null"`
	// 38	dealimuser	varchar2	8	yes
	DealimUser string `orm:"column(dealimuser);size(8);null"`
	// 39	validinfo	varchar2	60	yes
	ValidInfo string `orm:"column(validinfo);size(60);null"`
	// 40	pretranflag	varchar2	1	no
	// 默认值：'0'
	Pretranflag string `orm:"column(pretranflag);size(1);null"`

	//落地状态：71-落地拒绝；72-落地审批通过
	Down_stt string `orm:"column(down_stt);size(2);null"`

	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}

func (n *BaseSTMQudaoTranflow) TableName() string {
	return beego.AppConfig.String("base_stmqudaotranflow_table")
}

//验证用户信息
func checkBaseSTMQudaoTranflow(u *BaseSTMQudaoTranflow) (err error) {
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
	orm.RegisterModel(new(BaseSTMQudaoTranflow))
}

//get node list
func GetBaseSTMQudaoTranflowlist(page int64, page_size int64, sort string) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTranflow)
	qs := o.QueryTable(node)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&nodes, "Id", "Bizdate", "State", "Status")
	count, _ = qs.Count()
	return nodes, count
}

func GetBaseSTMQudaoTranflowById(nid int64) (BaseSTMQudaoTranflow, error) {
	o := orm.NewOrm()
	node := BaseSTMQudaoTranflow{Id: nid}
	err := o.Read(&node)
	if err != nil {
		return node, err
	}
	return node, nil
}

//添加用户
func AddBaseSTMQudaoTranflow(n *BaseSTMQudaoTranflow) (int64, error) {
	if err := checkBaseSTMQudaoTranflow(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTranflow)
	// node.Name = n.Name
	// node.State = n.State
	// node.Country = n.Country

	id, err := o.Insert(node)
	return id, err
}

//更新用户
func UpdateBaseSTMQudaoTranflow(n *BaseSTMQudaoTranflow) (int64, error) {
	if err := checkBaseSTMQudaoTranflow(n); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	node := make(orm.Params)
	//@todo 校验条件
	// if len(node) == 0 {
	// 	return 0, errors.New("update field is empty")
	// }
	var table BaseSTMQudaoTranflow
	num, err := o.QueryTable(table).Filter("Id", n.Id).Update(node)
	return num, err
}

func DelBaseSTMQudaoTranflowById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&BaseSTMQudaoTranflow{Id: Id})
	return status, err
}

func GetBaseSTMQudaoTranflowlistByGroupid(Groupid int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTranflow)
	count, _ = o.QueryTable(node).Filter("Group", Groupid).RelatedSel().Values(&nodes)
	return nodes, count
}

func GetBaseSTMQudaoTranflowTree(pid int64, level int64) ([]orm.Params, error) {
	o := orm.NewOrm()
	node := new(BaseSTMQudaoTranflow)
	var nodes []orm.Params
	_, err := o.QueryTable(node).Filter("Pid", pid).Filter("Level", level).Filter("Status", 2).Values(&nodes)
	if err != nil {
		return nodes, err
	}
	return nodes, nil
}

// GetAllDownloadtrac retrieves all Downloadtrac matches certain condition. Returns empty list if
// no records exist
func GetAllBaseSTMQudaoTranflows(query map[string]string, fields []string, sortby []string, order []string,
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
