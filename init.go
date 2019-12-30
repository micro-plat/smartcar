package main

import (
	"github.com/micro-plat/hydra/component"
)

//init 检查应用程序配置文件，并根据配置初始化服务
func init() {
	app.Initializing(func(c component.IContainer) error {
		return nil
	})
}
