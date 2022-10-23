package demo

import (
	"fmt"
	"time"

	"github.com/gohade/hade/framework/contract"
	"github.com/gohade/hade/framework/gin"
	"github.com/gohade/hade/framework/provider/redis"
)

// DemoRedis redis的路由方法
func (api *DemoApi) DemoRedis(c *gin.Context) {

	// 初始化一个redis
	redisService := c.MustMake(contract.RedisKey).(contract.RedisService)
	client, err := redisService.GetClient(redis.WithConfigPath("cache.default"), redis.WithRedisConfig(func(options *contract.RedisConfig) {
		options.MaxRetries = 3
	}))
	if err != nil {
		c.AbortWithError(50001, err)
		return
	}
	if err := client.Set(c, "foo", "bar", 1*time.Hour).Err(); err != nil {
		c.AbortWithError(500, err)
		return
	}
	val := client.Get(c, "foo").String()
	fmt.Println(val)

	if err := client.Del(c, "foo").Err(); err != nil {
		c.AbortWithError(500, err)
		return
	}

	c.JSON(200, "ok")
}

