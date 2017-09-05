package routers

import (
	"github.com/devuser/beego-admin-demo/models"
)

type ProductsRouter struct {
	baseRouter
}

func (this *ProductsRouter) Get() {
	this.TplName = "products.html"
	this.Data["IsProducts"] = true
	this.Data["Products"] = models.Products
}
