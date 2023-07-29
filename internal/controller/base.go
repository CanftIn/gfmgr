package controller

import (
	commonController "github.com/CanftIn/gfmgr/internal/common/controller"
	"github.com/gogf/gf/v2/net/ghttp"
)

type BaseController struct {
	commonController.BaseController
}

func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}
