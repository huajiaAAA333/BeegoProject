package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"io/ioutil"
)

type registerController struct {
	//beego..Controller
}

//func (r*registerController) Post {
	//解析前段提交的用户注册星系
//	dataBytes,err :=ioutil.ReadAll(r.Ctx.Request.Body)
	//if err != nil {
		r.Ctx.WriteString("数据接受错误")
		return

	}
	//var user models.User
	//err = json.Unmarshal(dataBytes,&user)
	//if err != nil {
		r.Ctx.WriteString("数据解析错误")
		return

	}
	//一切正常,将用户星系保存到数据库中取
	//直接调用保存数据的一个函数

	db_mysql.AddUser()
}


