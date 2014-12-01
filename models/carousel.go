package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

//首页轮换
type Carousel struct {
	Id         int64
	Title      string `orm:"size(100)"`
	Des        string `orm:"size(700)"`
	Picture    string `orm:"size(300)"`
	Url        string `"orm:size(300)"`
	UrlDes     string `"orm:size(200)"`
	Alt        string `"orm:size(200)"`
	Rank       int8
	Isactivity int8
	CreateTime time.Time `orm:"type(datetime)"`
	CreateUser string    `"orm:size(200)"`
	Ishide     int8
}

func (m *Carousel) TableName() string {
	return TableName("carousel")
}

func (m *Carousel) Insert() error {
	if _, err := orm.NewOrm().Insert(m); err != nil {
		return err
	}
	return nil
}

func (m *Carousel) Read(fields ...string) error {
	if err := orm.NewOrm().Read(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Carousel) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(m, fields...); err != nil {
		return err
	}
	return nil
}

func (m *Carousel) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func (m *Carousel) Query() orm.QuerySeter {
	return orm.NewOrm().QueryTable(m)
}
