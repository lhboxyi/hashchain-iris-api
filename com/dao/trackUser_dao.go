package dao

import (
	"github.com/go-xorm/xorm"
	"iris-api/com/models"
)



type TrackUserDao struct {
	engine *xorm.Engine
}


func NewTrackUserDao(engine *xorm.Engine) *TrackUserDao{
	return &TrackUserDao{
		engine:engine,
	}
}

/*func (tud *TrackUserDao) QueryUserByString(query string) (trackUser models.TrackUser) {

	var user models.TrackUser

	session := tud.engine.Table("track_user")
	if query!=""{
		session.Where(query)
	}

	err := session.Find(&user);

	if err != nil {
		return  user
	} else {
		return user
	}
}*/

//获取分页
func (d *TrackUserDao) QueryUserByString(query, sort string, pageSize int) []models.TrackUser{
	datalist := make([]models.TrackUser, 0)

	session := d.engine.Table("track_user")
	if query != "" {
		session.Where(query)
	}
	if sort != "" {
		session.OrderBy(sort)
	}
	if pageSize > 0 {
		limit := pageSize
		start := 0
		session.Limit(limit, start)
	}
	err := session.Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}

//获取分页
func (d *TrackUserDao) Login(userName string) []models.TrackUser{
	datalist := make([]models.TrackUser, 0)

	session := d.engine.Table("track_user")
	if userName != "" {
		session.Where("user_name=?",userName)
	}
	err := session.Find(&datalist)
	if err != nil {
		return datalist
	} else {
		return datalist
	}
}
