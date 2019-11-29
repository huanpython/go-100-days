/*
@Time : 2019/11/29 0029 下午 2:31
@Author : huanfuan
@File : user
@Software: GoLand
*/

package user

import "github.com/gogf/gf/frame/g"

const (
	USER_SESSION_MARK = "user_info"
)

var (
	// 表对象
	table = g.DB().Table("user").Safe()
)
