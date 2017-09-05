package models

import (
	"time"
)

type SQLModel struct {
	Id int64 `orm:"column(id);pk;auto"`

	Status uint64 `orm:"column(status);default(2)" form:"Status"`

	Creator  *User     `orm:"column(creator);null;rel(fk)"`
	CreateAt time.Time `orm:"column(create_at);type(datetime);auto_now_add"`

	Updater  *User     `orm:"column(updater);null;rel(fk)"`
	UpdateAt time.Time `orm:"column(update_at);type(datetime);null"`
	Message  string    `orm:"column(message);size(255);null"`
}
