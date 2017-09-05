package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
	"log"
	"time"
)

//角色表
type Role struct {
	Id    int64  `orm:"column(id);pk;auto"`
	Title string `orm:"size(255);unique;" form:"Title" valid:"Required"`
	Name  string `orm:"size(255);unique;" form:"Name" valid:"Required"`

	//是否管理员
	IsAdmin bool `orm:"default(0);null;column(is_admin)" form:"IsAdmin"`
	//是否安全审计人员
	IsAduditor bool `orm:"default(0);null;column(is_auditor)" form:"IsAduditor"`

	Remark string `orm:"null;size(200)" form:"Remark" valid:"MaxSize(200)"`
	//管控部分
	Status   uint64    `orm:"column(status);default(2)" form:"Status"`
	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`
	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`

	Node []*Node `orm:"reverse(many)"`
	User []*User `orm:"reverse(many)"`
}

func (r *Role) TableName() string {
	return beego.AppConfig.String("rbac_role_table")
}

func init() {
	orm.RegisterModel(new(Role))
}

func checkRole(g *Role) (err error) {
	valid := validation.Validation{}
	b, _ := valid.Valid(&g)
	if !b {
		for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			return errors.New(err.Message)
		}
	}
	return nil
}

//get role list
func GetRolelist(page int64, page_size int64, sort string) (roles []orm.Params, count int64) {
	o := orm.NewOrm()
	role := new(Role)
	qs := o.QueryTable(role)
	var offset int64
	if page <= 1 {
		offset = 0
	} else {
		offset = (page - 1) * page_size
	}
	qs.Limit(page_size, offset).OrderBy(sort).Values(&roles)
	count, _ = qs.Count()
	return roles, count
}

func AddRole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := new(Role)
	role.Title = r.Title
	role.Name = r.Name
	role.Remark = r.Remark
	role.Status = r.Status

	id, err := o.Insert(role)
	return id, err
}

func UpdateRole(r *Role) (int64, error) {
	if err := checkRole(r); err != nil {
		return 0, err
	}
	o := orm.NewOrm()
	role := make(orm.Params)
	if len(r.Title) > 0 {
		role["Title"] = r.Title
	}
	if len(r.Name) > 0 {
		role["Name"] = r.Name
	}
	if len(r.Remark) > 0 {
		role["Remark"] = r.Remark
	}
	if r.Status != 0 {
		role["Status"] = r.Status
	}
	if len(role) == 0 {
		return 0, errors.New("update field is empty")
	}
	var table Role
	num, err := o.QueryTable(table).Filter("Id", r.Id).Update(role)
	return num, err
}

func DelRoleById(Id int64) (int64, error) {
	o := orm.NewOrm()
	status, err := o.Delete(&Role{Id: Id})
	return status, err
}

func GetNodelistByRoleId(Id int64) (nodes []orm.Params, count int64) {
	o := orm.NewOrm()
	node := new(Node)
	count, _ = o.QueryTable(node).Filter("Role__Role__Id", Id).Values(&nodes)
	return nodes, count
}

func DelGroupNode(roleid int64, groupid int64) error {
	var nodes []*Node
	var node Node
	role := Role{Id: roleid}
	o := orm.NewOrm()
	num, err := o.QueryTable(node).Filter("Group", groupid).RelatedSel().All(&nodes)
	if err != nil {
		return err
	}
	if num < 1 {
		return nil
	}
	for _, n := range nodes {
		m2m := o.QueryM2M(n, "Role")
		_, err1 := m2m.Remove(&role)
		if err1 != nil {
			return err1
		}
	}
	return nil
}
func AddRoleNode(roleid int64, nodeid int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleid}
	node := Node{Id: (nodeid)}
	m2m := o.QueryM2M(&node, "Role")
	num, err := m2m.Add(&role)
	return num, err
}

func DelUserRole(roleid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("user_roles").Filter("role_id", roleid).Delete()
	return err
}
func AddRoleUser(roleid int64, userid int64) (int64, error) {
	o := orm.NewOrm()
	role := Role{Id: roleid}
	user := User{Id: (userid)}
	m2m := o.QueryM2M(&user, "Role")
	num, err := m2m.Add(&role)
	return num, err
}

func GetUserByRoleId(roleid int64) (users []orm.Params, count int64) {
	o := orm.NewOrm()
	user := new(User)
	count, _ = o.QueryTable(user).Filter("Role__Role__Id", roleid).Values(&users)
	return users, count
}

// func AccessList(uid int64) (list []orm.Params, err error) {
// 	var roles []orm.Params
// 	o := orm.NewOrm()
// 	role := new(Role)
// 	_, err = o.QueryTable(role).Filter("User__User__Id", uid).Values(&roles)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var nodes []orm.Params
// 	node := new(Node)
// 	for _, r := range roles {
// 		_, err := o.QueryTable(node).Filter("Role__Role__Id", r["Id"]).Values(&nodes)
// 		if err != nil {
// 			return nil, err
// 		}
// 		for _, n := range nodes {
// 			list = append(list, n)
// 		}
// 	}
// 	return list, nil
// }
//

func AccessList(uid int64) (list []orm.Params, err error) {
	var roles []orm.Params
	o := orm.NewOrm()
	role := new(Role)
	fmt.Printf("uid: %q\n", uid)
	//表名user
	//_, err = o.QueryTable(role).Filter("User__User__Id", uid).Values(&roles)
	//表名rbac_user
	//_, err = o.QueryTable(role).Values(&roles)

	//fmt.Printf("%q\n", roles[0])
	//panic(0)
	//@todo: 待核实
	_, err = o.QueryTable(role).Filter("User__Id", uid).Values(&roles)
	if err != nil {
		panic(err)
		return nil, err
	}
	var nodes []orm.Params
	node := new(Node)
	fmt.Printf("roles: %q\n", roles)
	for _, r := range roles {
		_, err := o.QueryTable(node).Filter("Role__Role__Id", r["Id"]).Values(&nodes)
		if err != nil {
			panic(err)
			return nil, err
		}
		fmt.Printf("nodes: %q\n", nodes)
		for _, n := range nodes {
			list = append(list, n)
		}
	}
	return list, nil
}
