/*
@Time : 2019/11/29 0029 下午 2:21
@Author : huanfuan
@File : chat
@Software: GoLand
*/

package chat

import (
	"github.com/gogf/gf/frame/gmvc"
	"github.com/gogf/gf/net/ghttp"
)

// Controller 控制器结构体
type Controller struct {
	gmvc.Controller
	ws *ghttp.WebSocket
}
