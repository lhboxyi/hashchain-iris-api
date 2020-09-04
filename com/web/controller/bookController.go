package controller

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"iris-api/com/service"
)

type BookController struct {
	Ctx     iris.Context
	Service service.IBookService
}


func NewBookController() *BookController {
	return &BookController{Service:service.NewBookService()}
}

func (c *BookController) BasicMVC(app *mvc.Application) {
	//当然，你可以在MVC应用程序中使用普通的中间件。
	app.Router.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("Path: %s", ctx.Path())
		ctx.Next()
	})
	//把依赖注入，controller(s)绑定
	//可以是一个接受iris.Context并返回单个值的函数（动态绑定）
	//或静态结构值（service）。
	/*app.Register(
		sessions.New(sessions.Config{}).Start,
		&prefixedLogger{prefix: "DEV"},
	)*/
	// GET: http://localhost:8080/basic
	// GET: http://localhost:8080/basic/custom
	app.Handle(new(BookController))
	//所有依赖项被绑定在父 *mvc.Application
	//被克隆到这个新子身上，父的也可以访问同一个会话。
	// GET: http://localhost:8080/basic/sub
	//app.Party("/sub").Handle(new(basicSubController))
}
func (c *BookController) Get() {
	service := service.NewBookService()
	list := service.GetList("", "ID asc", 0)
	/*return mvc.View{
		Name: "book/home.html",
		Data: iris.Map{
			"Title":  "首页-" ,
			"List":   list,
		},
		Layout: "shared/bookLayout.html",
	}*/
	c.Ctx.JSON(list)
}

///book/ajaxbooks?key=go	访问地址是小写的
func (c *BookController) GetAjaxbooks() {
	//获取url参数
	key := c.Ctx.URLParam("key")

	service := service.NewBookService()
	list := service.GetList(" bookName like '%"+key+"%'", "ID asc", 0)

	c.Ctx.JSON(list)
}