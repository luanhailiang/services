package handler

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"

	"github.com/luanhailiang/micro/plugins/matedata"
	rds "github.com/luanhailiang/micro/plugins/store/redis"
	online "github.com/luanhailiang/services/online/proto"
)

type GrpcHandler struct {
}

var keys map[string]bool

func (p *GrpcHandler) CmdHeart(ctx context.Context, cmd *online.CmdHeart) *online.CmdHeart {
	mate, _ := matedata.FromMateContext(ctx)
	m := "ios"
	if cmd.Plat == 1 {
		m = "and"
	}
	key := fmt.Sprintf("beat_%s_%s", mate.Space, m)
	key = rds.Key(ctx, key)
	client := rds.GetClient()
	client.ZAdd(ctx, key, &redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: mate.Index,
	})
	return cmd
}

func (p *GrpcHandler) CmdOnlineList(ctx context.Context, cmd *online.CmdOnlineList) (*online.CmdOnlineList, uint32, string) {
	mate, _ := matedata.FromMateContext(ctx)
	space := mate.Index
	if space == "" {
		space = cmd.Space
	}
	client := rds.GetClient()

	key_ios := fmt.Sprintf("online_beat_%s_ios", space)
	// err := client.ZRemRangeByScore(ctx, key_ios, "", fmt.Sprintf("%d", (time.Now().Unix()-300))).Err()
	// if err != nil {
	// 	logrus.Error("ZRemRangeByScore:%s:%s", key_ios, err.Error())
	// 	return cmd, uint32(codes.Unknown), err.Error()
	// }
	list_ios, err := client.ZRange(ctx, key_ios, 0, -1).Result()
	logrus.Debugf("CmdOnlineList:%s:%v", key_ios, list_ios)
	if err != nil {
		logrus.Error("CmdOnlineList:%s:%s", key_ios, err.Error())
		return cmd, uint32(codes.Unknown), err.Error()
	}
	key_and := fmt.Sprintf("online_beat_%s_and", space)
	// err = client.ZRemRangeByScore(ctx, key_and, "", fmt.Sprintf("%d", (time.Now().Unix()-300))).Err()
	// if err != nil {
	// 	logrus.Error("ZRemRangeByScore:%s:%s", key_and, err.Error())
	// 	return cmd, uint32(codes.Unknown), err.Error()
	// }
	list_and, err := client.ZRange(ctx, key_and, 0, -1).Result()
	logrus.Debugf("CmdOnlineList:%s:%v", key_and, list_and)
	if err != nil {
		logrus.Error("CmdOnlineList:%s:%s", key_and, err.Error())
		return cmd, uint32(codes.Unknown), err.Error()
	}
	cmd.List = append(cmd.List, list_ios...)
	cmd.List = append(cmd.List, list_and...)
	return cmd, 0, ""
}

func (p *GrpcHandler) CmdOnlineCount(ctx context.Context, cmd *online.CmdOnlineCount) (*online.CmdOnlineCount, uint32, string) {
	client := rds.GetClient()

	key_ios := fmt.Sprintf("online_beat_%d_ios", cmd.Space)
	num_ios, err := client.ZCard(ctx, key_ios).Result()
	if err != nil {
		logrus.Error("ZScore:%s:%s", key_ios, err.Error())
		return cmd, uint32(codes.Unknown), err.Error()
	}
	key_and := fmt.Sprintf("online_beat_%d_and", cmd.Space)
	num_and, err := client.ZCard(ctx, key_and).Result()
	if err != nil {
		logrus.Error("ZScore:%s:%s", key_and, err.Error())
		return cmd, uint32(codes.Unknown), err.Error()
	}
	cmd.Amount = uint64(num_and) + uint64(num_ios)
	return cmd, 0, ""
}

func (p *GrpcHandler) CmdOnlineAllCount(ctx context.Context, cmd *online.CmdOnlineAllCount) (*online.CmdOnlineAllCount, uint32, string) {
	datas := getCount()
	for _, data := range datas {
		cmd.List = append(cmd.List, data)
	}
	return cmd, 0, ""
}
