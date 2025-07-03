package global

import (
	"context"
	"fmt"
	"os"

	"github.com/redis/go-redis/v9"
)

var (
	Rdb *redis.Client
)

func InitRedis() error {
	dsn := os.Getenv("REDIS")

	if dsn == "" {
		return fmt.Errorf("REDIS: 环境变量未设置")
	}
	opt, err := redis.ParseURL(dsn)
	if err != nil {
		return err
	}

	Rdb = redis.NewClient(opt)

	ctx := context.Background()
	_, err = Rdb.Ping(ctx).Result()
	if err != nil {
		return fmt.Errorf("redis 连接失败: %s", err.Error())
	}

	fmt.Println("Redis 连接成功")
	return nil
}
