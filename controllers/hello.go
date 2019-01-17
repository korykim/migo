package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"migo/models"
)

// HelloController operations for Hello
type HelloController struct {
	beego.Controller
}

func add_news(){
 	o :=orm.NewOrm()
	news := new(models.News)
	news.Title="hello"
	news.Body="beego"
	fmt.Print(o.Insert(news))

}

func (c *HelloController) Get(){
	//add_news()
	c.Data["Name"] = "Migo"
	c.TplName="hello/index.tpl"
}

