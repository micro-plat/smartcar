package main

import (
	"github.com/micro-plat/hydra/context"
)

//bind 检查应用程序配置文件，并根据配置初始化服务
func init() {
	app.Handling(func(ctx *context.Context) (rt interface{}) {
		return nil
	})
}
