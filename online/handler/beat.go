package handler

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/luanhailiang/micro/plugins/events/nats_pub"
	rds "github.com/luanhailiang/micro/plugins/store/redis"
	"github.com/luanhailiang/micro/proto/events"
	"github.com/luanhailiang/micro/proto/rpcmsg"
	online "github.com/luanhailiang/services/online/proto"
	"github.com/sirupsen/logrus"
)

// 后期重写改成k8s定时任务
func init() {
	ticker := time.NewTicker(60 * time.Second)
	go func() {
		for t := range ticker.C {
			beatcheck(t)
		}
	}()
}

func getCount() map[int]*online.SpaceCount {
	client := rds.GetClient()
	ctx := context.Background()
	keys, err := client.Keys(ctx, "online_beat_*").Result()
	if err != nil {
		logrus.Error("online:%s", err.Error())
		return nil
	}
	datas := map[int]*online.SpaceCount{}
	for _, key := range keys {
		space := strings.Replace(key, "online_beat_", "", 1)

		args := strings.Split(space, "_")
		if len(args) != 2 {
			logrus.Error("len(args):%s", key)
			return nil
		}
		zone, err := strconv.Atoi(args[0])
		if err != nil {
			logrus.Error("strconv:%s:%s", space, err.Error())
			return nil
		}
		data, ok := datas[zone]
		if !ok {
			data = &online.SpaceCount{}
			data.Space = uint32(zone)
			datas[zone] = data
		}
		err = client.ZRemRangeByScore(ctx, key, "", fmt.Sprintf("%d", (time.Now().Unix()-300))).Err()
		if err != nil {
			logrus.Error("ZRemRangeByScore:%s:%s", key, err.Error())
			return nil
		}
		count, err := client.ZCard(ctx, key).Result()
		if err != nil {
			logrus.Error("ZScore:%s:%s", key, err.Error())
			return nil
		}
		if args[1] == "and" {
			data.AndCount = uint32(count)
		} else {
			data.IosCount = uint32(count)
		}
	}
	return datas
}

func beatcheck(t time.Time) {
	logrus.Debugf("online tick:%s", t.Local().String())
	datas := getCount()
	for zone, data := range datas {
		e := &events.OnlineCount{
			Time: t.Unix(),
			Game: "",
			Zone: int32(zone),
			Area: "中国",
			Ios:  int32(data.IosCount),
			And:  int32(data.AndCount),
		}
		nats_pub.PubMate(&rpcmsg.MateMessage{}, e)
	}
}
