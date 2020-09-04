package models

import (
	"time"
)

type TrackUser struct {
	Id int64 `xorm:int(64) id unsigned NOT NULL AUTO_INCREMENT`
	UserName string `xorm:varchar(128) user_name`
	Password string `xorm:varchar(256) password`
	CompanyName string `xorm: varchar(256) company_name`
	Represent string `xorm:varchar(256) represent`
	Account string `xorm:varchar(128) account`
	Contact string `xorm:varchar(40) contact`
	Phone string `xorm:varchar(40) phone`
	Address string `xorm:varchar(128) address`
	ZipCode string `xorm:varchar(40) zip_code`
	CreateTime time.Time `xorm:timestamp create_time`
	UpdateTime time.Time `xorm:timestamp update_time`
	LoginTime time.Time `xorm:timestamp login_time`
	Token      string `xorm:"-"`
	Session      string `xorm:"-"`
}
