package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/spf13/cast"
	"iris-api/com/service"
	"log"
)

type TrackUserController struct {
	Ctx     iris.Context
	Service service.ITrackUserService
}

func NewTrackUserController() *TrackUserController{
	return &TrackUserController{
		Service: service.NewTrackUserService(),
	}
}
func (c *TrackUserController) BasicMVC(app *mvc.Application) {
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})

	app.Handle(new(TrackUserController))

}

func (tuc *TrackUserController) PostLogin(){
	var m map[string]interface{}
	err := tuc.Ctx.ReadJSON(&m)
	if err != nil {
		log.Println("ReadJSON Error:", err)
	}
	//获取url参数
	userName := cast.ToString(m["userName"])

	password := cast.ToString(m["password"])

	service := service.NewTrackUserService();

	user := service.Login(userName,password)

	tuc.Ctx.JSON(user)

}


