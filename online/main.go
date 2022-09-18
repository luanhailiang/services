package main

import (
	"os"

	"net/http"
	_ "net/http/pprof"

	"github.com/sirupsen/logrus"

	"github.com/luanhailiang/micro/plugins/message/grpc_svr"
	"github.com/luanhailiang/services/online/handler"
)

func main() {
	if os.Getenv("DEBUG_LOG") != "" {
		logrus.SetLevel(logrus.DebugLevel)
	}
	go func() {
		logrus.Info(http.ListenAndServe(":6060", nil))
	}()

	m := &handler.GrpcHandler{}
	grpc_svr.Register(&m)

	grpc_svr.StartServe()
}
