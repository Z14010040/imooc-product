package contollers

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"imooc-product/services"
)

type Orderontroller struct{
	Ctx iris.Context
	OrderService services.OrderService
}
func (o *Orderontroller)Get()mvc.View{
	ordersInfo, err := o.OrderService.GetAllOrderInfo()
	if err!=nil{
		o.Ctx.Application().Logger().Debug("查询订单信息失败")
	}
	return mvc.View{
		Name:"order/view.html",
		Data:iris.Map{
			"order":ordersInfo,
		},
	}
}