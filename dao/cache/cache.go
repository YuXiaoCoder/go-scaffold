package cache

import (
	"context"
	"errors"
	"fmt"
	"go-scaffold/pkg/configs"
	"time"

	"github.com/go-redis/redis/v8"
)

// 全局变量
var (
	rdb *redis.Client
)

func Init() error {
	if configs.AllConfig.Cache.DriverName == "redis" {
		rdb = redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", configs.AllConfig.Cache.Redis.Host, configs.AllConfig.Cache.Redis.Port),
			Password: configs.AllConfig.Cache.Redis.Password,
			DB:       configs.AllConfig.Cache.Redis.DB,
			PoolSize: configs.AllConfig.Cache.Redis.PoolSize,
		})

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		// 测试连通性
		_, err := rdb.Ping(ctx).Result()
		if err != nil {
			return err
		}
	} else {
		return errors.New(fmt.Sprintf("unsupported cache driver name: %s", configs.AllConfig.Cache.DriverName))
	}

	return nil
}

func Close() {
	_ = rdb.Close()
}
