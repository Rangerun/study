package test

import (
	"fmt"

	"github.com/gohade/hade/framework"
)

// 服务提供方
type TestServiceProvider struct {
}

// Name方法直接将服务对应的字符串凭证返回，在这个例子中就是“hade.demo"
func (sp *TestServiceProvider) Name() string {
	return Key
}

// Register方法是注册初始化服务实例的方法，我们这里先暂定为NewDemoService
func (sp *TestServiceProvider) Register(c framework.Container) framework.NewInstance {
	return NewTestService
}

// IsDefer方法表示是否延迟实例化，我们这里设置为true，将这个服务的实例化延迟到第一次make的时候
func (sp *TestServiceProvider) IsDefer() bool {
	return true
}

// Params方法表示实例化的参数。我们这里只实例化一个参数：container，表示我们在NewDemoService这个函数中，只有一个参数，container
func (sp *TestServiceProvider) Params(c framework.Container) []interface{} {
	return []interface{}{c}
}

// Boot方法我们这里我们什么逻辑都不执行, 只打印一行日志信息
func (sp *TestServiceProvider) Boot(c framework.Container) error {
	fmt.Println("test service boot")
	return nil
}