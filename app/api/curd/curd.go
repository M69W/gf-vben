package curd

import (
	"Gf-Vben/app/service/curd"
	"Gf-Vben/app/service/response"
	"Gf-Vben/app/service/user"
	"github.com/gogf/gf/net/ghttp"
)

type Controller struct {
	Interface string `p:"i" v:"required"`
	Action    string `p:"a" v:"required"`
}

func (c *Controller) Curd(r *ghttp.Request) {
	if err := r.Parse(&c); err != nil {
		response.JsonExit(r, 1, err.Error())
	}
	var cu curd.Curd
	switch c.Interface {
	case "user":
		req := new(user.Req)
		cu = req
	default:
		response.JsonExit(r, 2, "接口参数错误")
	}
	if err := r.Parse(cu); err != nil {
		response.JsonExit(r, 3, err.Error())
	}
	switch c.Action {
	case "list":
		result, err := cu.List()
		if err != nil {
			response.JsonExit(r, 4, err.Error())
		}
		response.JsonExit(r, 0, "", result)
	case "add":
		if err := cu.Add(); err != nil {
			response.JsonExit(r, 4, err.Error())
		}
		response.JsonExit(r, 0, "新增成功")
	case "edit":
		if err := cu.Edit(); err != nil {
			response.JsonExit(r, 4, err.Error())
		}
		response.JsonExit(r, 0, "修改成功")
	case "del":
		if err := cu.Del(); err != nil {
			response.JsonExit(r, 4, err.Error())
		}
		response.JsonExit(r, 0, "删除成功")
	default:
		response.JsonExit(r, 4, "接口参数错误")
	}
}
