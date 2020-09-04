package models

import (
	"time"
)

type BookTb struct {
	Id           int       `xorm:"not null pk autoincr INT(11)"`
	BookName     string    `xorm:"default '' comment('书名') VARCHAR(100)"`
	State        int       `xorm:"default 0 comment('状态 0:未出借,1:出借') INT(11)"`
	Author       string    `xorm:"default '' comment('作者') VARCHAR(50)"`
	Press        string    `xorm:"default '' comment('出版社') VARCHAR(100)"`
	PublishTime  time.Time `xorm:"comment('出版时间') DATETIME"`
	BookImage    string    `xorm:"default '' comment('图书封面') VARCHAR(300)"`
	Price        string    `xorm:"default 0.00 comment('售价') DECIMAL(10,2)"`
	Introduction string    `xorm:"default '' comment('简介') VARCHAR(300)"`
	UpdateTime   time.Time `xorm:"DATETIME"`
	AddTime      time.Time `xorm:"DATETIME"`
}