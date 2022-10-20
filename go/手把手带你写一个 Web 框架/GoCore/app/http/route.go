package http

import (
	"github.com/gohade/hade/app/http/middleware/cors"
	"github.com/gohade/hade/app/http/module/demo"
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/framework/middleware/static"
	"github.com/gohade/hade/framework"
)

// NewHttpEngine 创建了一个绑定了路由的Web引擎
func NewHttpEngine(container framework.Container) (*gin.Engine, error) {
	// 设置为Release，为的是默认在启动中不输出调试信息
	gin.SetMode(gin.ReleaseMode)
	// 默认启动一个Web引擎
	r := gin.New()
	// 设置了Engine
	r.SetContainer(container)

	// 默认注册recovery中间件
	r.Use(gin.Recovery())

	// 业务绑定路由操作
	Routes(r)
	// 返回绑定路由后的Web引擎
	return r, nil
}


// Routes 绑定业务层路由
func Routes(r *gin.Engine) {
	
	// /路径先去./dist目录下查找文件是否存在，找到使用文件服务提供服务
	r.Use(static.Serve("/", static.LocalFile("./dist", false)))
	// 使用cors中间件
	r.Use(cors.Default())

	

	// 动态路由定义
	demo.Register(r)
}
