package service

import (
	"golang.org/x/crypto/bcrypt"
	"iris-api/com/dao"
	datasource "iris-api/com/dbsource"
	"iris-api/com/middleware"
	"iris-api/com/models"
)

type ITrackUserService interface {
	Login(userName string,password string) (result models.Result)
}

type trackUserService struct {
	dao *dao.TrackUserDao
}

func NewTrackUserService() ITrackUserService{
	return &trackUserService{
		dao : dao.NewTrackUserDao(datasource.InstanceMaster()),
	}
}



func (tus *trackUserService) Login(userName string,password string) (result models.Result){
	if userName == "" {
		result.Success = "false"
		result.Code = "-1"
		result.Msg = "请输入用户名！"
		return
	}
	if password == "" {
		result.Success = "false"
		result.Code = "-1"
		result.Msg = "请输入密码！"
		return
	}

	var user = tus.dao.QueryUserByString("user_name='"+userName+"'","create_time",1);

	if len(user) == 0{
		result.Success = "false"
		result.Code = "-1"
		result.Msg = "用户不存在！"
		return
	}
	// 正确密码验证
	err := bcrypt.CompareHashAndPassword([]byte(user[0].Password), []byte(password))
	if err != nil {
		result.Success = "false"
		result.Code = "-1"
		result.Msg = "密码错误！"
		return
	} else {
		/*passwordEncoder,err:= bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Println(err)
		}*/
		userOne := tus.dao.Login(userName)[0];
		if len(user) == 0 {
			result.Success = "false"
			result.Code = "-1"
			result.Msg = "用户名或密码错误!"
			return
		}
		userOne.Token = middleware.GenerateToken(userOne)
		result.Success = "true"
		result.Code = "0"
		result.Data = userOne
		result.Msg = "登录成功"
		return
	}

}