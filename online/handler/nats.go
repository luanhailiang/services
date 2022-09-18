package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/luanhailiang/micro/plugins/events/nats_sub"
	"github.com/luanhailiang/micro/plugins/matedata"
	rds "github.com/luanhailiang/micro/plugins/store/redis"
	"github.com/luanhailiang/micro/proto/events"
)

func init() {
	m := &NatsHandler{}
	nats_sub.Register(&m)
}

type NatsHandler struct {
}

func (p *NatsHandler) EVENTS__SignIn(ctx context.Context, event *events.SignIn) {
	mate, _ := matedata.FromMateContext(ctx)
	m := "ios"
	if event.Plat == 1 {
		m = "and"
	}
	key := fmt.Sprintf("beat_%s_%s", mate.Space, m)
	key = rds.Key(ctx, key)
	client := rds.GetClient()
	client.ZAdd(ctx, key, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: mate.Index,
	})
	client.ZIncrBy(ctx, rds.Key(ctx, "count_"+mate.Space), 1, m)
}

func (p *NatsHandler) EVENTS__Heart(ctx context.Context, event *events.Heart) {
	mate, _ := matedata.FromMateContext(ctx)
	m := "ios"
	if event.Plat == 1 {
		m = "and"
	}
	key := fmt.Sprintf("beat_%s_%s", mate.Space, m)
	key = rds.Key(ctx, key)
	client := rds.GetClient()
	client.ZAdd(ctx, key, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: mate.Index,
	})
}

func (p *NatsHandler) EVENTS__Logout(ctx context.Context, event *events.Logout) {
	mate, _ := matedata.FromMateContext(ctx)
	m := "ios"
	if event.Plat == 1 {
		m = "and"
	}
	key := fmt.Sprintf("beat_%s_%s", mate.Space, m)
	key = rds.Key(ctx, key)
	client := rds.GetClient()
	client.ZRem(ctx, key, mate.Index)
	client.ZIncrBy(ctx, rds.Key(ctx, "count_"+mate.Space), -1, m)
}
