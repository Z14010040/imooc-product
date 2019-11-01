package main

import (
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imooc-product/backend/web/contollers"
	"imooc-product/common"
	"imooc-product/repositories"
	"imooc-product/services"
	"log"
)

func main()  {
	//1、创建iris实例
	app:=iris.New()

	//2、设置错误模式，在mvc模式下提示错误
	app.Logger().SetLevel("debug")

	//3、注册模板  用来渲染视图
	template:=iris.HTML("./backend/web/views",".html").
		//设置布局文件
		Layout("shared/layout.html").Reload(true)
	app.RegisterView(template)

	//4、设置模板目标
	app.HandleDir("/assets","./backend/web/assets")
	//出现异常跳转到异常页面
	app.OnAnyErrorCode(func(context iris.Context) {
		context.ViewData("message",context.Values().GetStringDefault("message","访问的页面出错"))
		context.ViewLayout("")
		context.View("shared/error.html")
	})
	
	//连接数据库
	db,err:=common.NewMysqlConn()
	if err!=nil{
		log.Println("[*] waiting for messages, To exit press CTRL+C")
	}
	//上下文对象
	ctx,cancel:=context.WithCancel(context.Background())
	defer cancel()
	
	//5、注册控制器
	productRepository:=repositories.NewProductRepositoryManager("product",db)
	productService:=services.NewProductServiceManager(productRepository)
	productParty:=app.Party("/product")
	product:=mvc.New(productParty)
	product.Register(ctx,productService)
	product.Handle(new(contollers.ProductController))
	
	orderRepository:=repositories.NewOrderRepositoryManager("order",db)
	orderService:=services.NewOrderServiceManager(orderRepository)
	orderParty:=app.Party("/order")
	order:=mvc.New(orderParty)
	order.Register(ctx,orderService)
	order.Handle(new(contollers.Orderontroller))

	//6.启动服务器
	app.Run(
		iris.Addr("localhost:8080"),
		iris.WithoutServerError(iris.ErrServerClosed),
		iris.WithOptimizations,)
}