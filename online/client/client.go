package online_cli

import (
	"context"
	"fmt"
	"os"
	"strconv"

	"github.com/luanhailiang/micro/plugins/message/grpc_cli"
	mem "github.com/luanhailiang/micro/plugins/store/cache"
	online "github.com/luanhailiang/services/online/proto"
	"github.com/sirupsen/logrus"
	"google.golang.org/protobuf/proto"
)

var localCache *mem.Manager

func init() {
	if os.Getenv("CLIENT_CACHE") != "" {
		localCache = mem.NewManager()

		over, err := strconv.Atoi(os.Getenv("CLIENT_CACHE"))
		if err != nil {
			logrus.Fatalf("CLIENT_CACHE:%s %s", os.Getenv("CLIENT_CACHE"), err.Error())
		}
		localCache.SetOver(int64(over))
	}
}

func OnlineAllCount(ctx context.Context) (*online.CmdOnlineAllCount, error) {
	cmd := &online.CmdOnlineAllCount{}
	ret, err := grpc_cli.Call(ctx, cmd)
	if err != nil {
		return nil, err
	}
	if ret.Code != 0 {
		return nil, fmt.Errorf(ret.Info)
	}
	err = proto.Unmarshal(ret.Buff.Data, cmd)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func OnlineCount(ctx context.Context, cmd *online.CmdOnlineCount) error {
	index := fmt.Sprintf("%d", cmd.Space)
	if localCache != nil {
		err := localCache.GetOne(ctx, index, cmd)
		if err == nil {
			return nil
		}
	}
	ret, err := grpc_cli.Call(ctx, cmd)
	if err != nil {
		return err
	}
	if ret.Code != 0 {
		return fmt.Errorf(ret.Info)
	}
	err = proto.Unmarshal(ret.Buff.Data, cmd)
	if err != nil {
		return err
	}
	localCache.SetOne(ctx, index, cmd)
	return nil
}

func OnlineList(ctx context.Context, space string) (*online.CmdOnlineList, error) {
	cmd := &online.CmdOnlineList{
		Space: space,
	}
	ret, err := grpc_cli.Call(ctx, cmd)
	if err != nil {
		return nil, err
	}
	if ret.Code != 0 {
		return nil, fmt.Errorf(ret.Info)
	}
	err = proto.Unmarshal(ret.Buff.Data, cmd)
	if err != nil {
		return nil, err
	}
	return cmd, err
}
