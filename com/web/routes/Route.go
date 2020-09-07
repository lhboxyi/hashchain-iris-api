package routes

import (
	"encoding/json"
	"fmt"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io/ioutil"
	"iris-api/com/middleware"
	"iris-api/com/util"
	"iris-api/com/web/controller"
	"net/http"
	"strings"
	"time"
)



func Register() {
	//获取服务主机端口配置信息
	appPath := viper.Get("appserver.path").(string)
	appPort := viper.Get("appserver.port").(int)
	outFileDir := viper.Get("openapi.outFileDir").(string)
	outFileName := viper.Get("openapi.outFileName").(string)

	app := iris.New()
	// 注册 "before"  处理器作为当前域名所有路由中第一个处理函数
	app.Use(before)
	// 主持后置
	app.DoneGlobal(after)

	crs := cors.New(cors.Options{
		// allows everything, use that to change the hosts.
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
		AllowCredentials: true,
	})

	//未登录可以访问的接口
	mvc.Configure(
		app.Party("/a",crs),
		controller.NewTrackUserController().BasicMVC)
	app.Use(middleware.GetJWT().Serve)  // jwt
	//登录可以访问的接口
	mvc.Configure(
		app.Party("/basic",crs),
		controller.NewBookController().BasicMVC)

	//初始化中间件
	/*yaag.Init(&yaag.Config{
		On:       true, //是否开启自动生成API文档功能
		DocTitle: "Iris",
		DocPath:  outFileDir + "/" + outFileName, //生成API文档名称存放路径
		BaseUrls: map[string]string{"Production": "", "Staging": ""},
	})*/
	//注册中间件
	//app.Use(irisyaag.New())



	//作用范围 全局 app.Use(authentication) 或者 (app.UseGlobal 在Run之前)
	//作用范围 单个路由 app.Get("/mysecret", authentication, h)
	app.Get("/", func(ctx iris.Context) { ctx.Redirect("/admin") })
	//作用范围  Party

	//开放OPEN API FILE
	openApi(app, appPath, outFileDir, outFileName)




	/*app.Run(iris.Addr(fmt.Sprintf(":%d", appPort)),
		iris.WithoutServerError(iris.ErrServerClosed),//忽略服务器错误
		iris.WithOptimizations,//让程序自身尽可能的优化
		iris.WithCharset("UTF-8"), // 国际化
	)*/

	app.Run(iris.TLS(fmt.Sprintf(":%d", appPort),"/home/zhangzhemin/iris-api/ellipticcurve_public.crt","/home/zhangzhemin/iris-api/ellipticcurve.key"),
	)

	//当报 `404` 时候渲染自定义的 404 错误模板
	app.OnErrorCode(iris.StatusNotFound, func(ctx iris.Context){
		ctx.View("errors/404.html")
	})
	//app.Handle("","",controller.NewBookController())
	app.OnErrorCode(500, func(ctx iris.Context){
		// 编写对500 错误的响应
	})
}




/**
开放接口描述文档
*/
func openApi(crawlApp *iris.Application, appPath, outFileDir, outFileName string) {
	crawlApp.RegisterView(iris.HTML("view", ".html"))
	crawlApp.Handle("GET", appPath+"/openApi", func(ctx iris.Context) {
		path := outFileDir + "/" + outFileName
		html, _ := util.ReadAll(path)
		ctx.HTML(html)
	})
}

/**
前置函数"before"  处理器作为当前域名所有路由中第一个处理函数
获取每次请求访问前的接口路径、请求接口的ip、请求时间
*/
func before(ctx iris.Context) {
	requestPath := ctx.Path()
	requestMethod := ctx.Request().Method
	clientAddr := strings.Split(ctx.Request().RemoteAddr, ":")
	requestDate := util.GetUnixToFormatString(time.Now().Unix(), "2006-01-02 15:04:05")
	logrus.Infof("请求接口方式【%s】,请求路径【%s】,请求客户端ip【%s】,请求时间【%s】", requestMethod, requestPath, clientAddr[0], requestDate)
	// 执行下一个处理器
	ctx.Next()
}

func after(ctx iris.Context) {
	ctx.WriteString("结束时间：" + util.GetUnixToFormatString(time.Now().Unix(), "2006-01-02 15:04:05"))
}

/**
验证请求头中的token
*/
func validateToken(ctx iris.Context) {
	//判断请求类型，如果是OPTIONS请求形式，则放行，其他请求均需验证token
	methodType := ctx.Request().Method
	if methodType == "OPTIONS" {
		ctx.Next()
	} else {
		tokenStr := ctx.GetHeader("Authorization")
		tokenUrl := viper.GetString("validatetokenurl")
		resp, err := http.Get(tokenUrl + tokenStr)
		if err != nil {
			ctx.JSON(err)
		}
		bytes, _ := ioutil.ReadAll(resp.Body)
		resMap := make(map[string]string)
		json.Unmarshal(bytes, &resMap)
		if resMap["responseCode"] == "S200" {
			// 执行下一个处理器
			ctx.Next()
		} else {
			result := controller.InstanceResult(resMap["responseCode"], resMap["data"], resMap["message"])
			ctx.JSON(result)
		}
	}
}

func h(ctx iris.Context) {
	username, password, _ := ctx.Request().BasicAuth()
	//第三个参数因为中间件所以不需要判断其值，否则不会执行此处理程序
	ctx.Writef("%s %s:%s", ctx.Path(), username, password)
}


