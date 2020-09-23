package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller //匿名字段
}

func (c *MainController) Get() {
	//获取get请求的不同方法
	name := c.Ctx.Input.Query("name")
	age := c.Ctx.Input.Query("age")
	sex := c.Ctx.Input.Query("sex")
	fmt.Println(name,age,sex)
	//以dong,8为正确数据进行验证
	if name != "dong"||age != "18"||sex != "female"{
		 c.Ctx.ResponseWriter.Write([]byte("数据验证成功"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("数据验证失败"))




	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "2509802388@qq.com"
	c.TplName = "index.tpl"
}
/**
该post方法是处理post类型的的请求的时候要调用的方法
 */
func (c * MainController) Post(){
	fmt.Println("post类型的请求")
	user:= c.Ctx.Request.FormValue("user")
	fmt.Println("用户名为:",user)
	psd := c.Ctx.Request.FormValue("psd")
	fmt.Println("密码是:",psd)

	//与固定值比较
	if user != "dong" || psd != "123456"{
		//失败页面
		c.Ctx.ResponseWriter.Write([]byte("对不起,数据不正确"))
		return
	}
	c.Ctx.ResponseWriter.Write([]byte("恭喜你,数据正确"))

	//request 请求, response响应


}