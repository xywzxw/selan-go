package controllers

import (
	"fmt"
	"selan/dbhelper"
	"selan/utils"
)

type baseparam struct {
	Act string `form:"act"`
	Op  string `form:"op"`
}

type SeLanController struct {
	BaseController
}

func (vc *SeLanController) Get() {
	vc.TplName = "outpage.tpl"

	u := baseparam{}
	if err := vc.ParseForm(&u); err != nil {
		//handle error
		fmt.Println("数据错误")
	}
	if u.Act == "" || u.Op == "" {
		vc.Ctx.WriteString("参数错误，缺少act,op参数 ")
		return
	}

	if u.Op == "area" {
		sqlStr := "select * from slm_area limit 10"
		vc.Ctx.WriteString(dbhelper.SQLJSONData(sqlStr))
	} else if u.Op == "map" {
		// 直接创建
		map1 := make(map[string]interface{})
		// 然后赋值
		map1["code"] = 200

		map2 := make(map[string]interface{})
		map2["result"] = "success"
		map2["desc"] = "成功"
		map1["data"] = map2

		array := []string{"123", "345", "ddsdd"}
		map1["arr"] = array

		vc.Ctx.WriteString(utils.ChangeMapToJSON(map1))
	} else if u.Op == "head" {
		// v := vc.Ctx.Request.URL
		// vc.Ctx.WriteString(v)
	} else if u.Op == "filter" {

	}
}

func (vc *SeLanController) Post() {
	vc.TplName = "outpage.tpl"

	act := vc.GetString("act")
	// op := vc.GetString("op")

	u := baseparam{}
	if err := vc.ParseForm(&u); err != nil {
		//handle error
	}
	vc.Ctx.WriteString("act:-----" + u.Act + "\n")
	vc.Ctx.WriteString("op:-----" + u.Op + "\n")

	vc.Ctx.WriteString("post请求" + act + "-------")
	if act == "" {
		vc.Ctx.WriteString("act is empty")
		return
	}
	if act == "index" {
		vc.Ctx.WriteString("进入首页")
	}

}
