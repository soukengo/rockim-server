package cache

import (
	"context"
	"fmt"
	"rockimserver/pkg/component/cache"
	rediscache "rockimserver/pkg/component/cache/redis"
	"rockimserver/pkg/component/database/redis"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	cli := redis.New(&redis.Config{Addr: []string{"127.0.0.1:6379"}})
	ch := rediscache.NewValueCache[User](cli, "user_test", cache.WithExpire(time.Second*10))
	ctx := context.TODO()
	val := &User{Id: "10001", Name: "张三", CreateTime: time.Now().UnixMilli()}
	err := ch.Set(ctx, val)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	v1, err := ch.Get(ctx)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	fmt.Printf("v1: %v \n", v1)
}
