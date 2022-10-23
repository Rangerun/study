package orm

import (
	"fmt"

	"github.com/gohade/hade/framework"
	"github.com/gohade/hade/framework/contract"

)

// GetBaseConfig 读取database.yaml根目录结构
func GetBaseConfig(c framework.Container) *contract.DBConfig {
	configService := c.MustMake(contract.ConfigKey).(contract.Config)
	config := &contract.DBConfig{}
	// 直接使用配置服务的load方法读取,yaml文件
	err := configService.Load("database", config)
	if err != nil {
		fmt.Println("配置出错")
		return nil
	}
	return config
}

// WithConfigPath 加载配置文件地址
func WithConfigPath(configPath string) contract.DBOption {
	return func(container framework.Container, config *contract.DBConfig) error {
		configService := container.MustMake(contract.ConfigKey).(contract.Config)
		// 加载configPath配置路径
		if err := configService.Load(configPath, config); err != nil {
			return err
		}
		return nil
	}
}
