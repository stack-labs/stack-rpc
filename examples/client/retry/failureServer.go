package main

import (
	"context"
	"time"

	"github.com/stack-labs/stack"
	proto "github.com/stack-labs/stack/examples/proto/service/rpc"
	log "github.com/stack-labs/stack/logger"
	"github.com/stack-labs/stack/util/errors"
)

type FailureExample struct{}

func (e *FailureExample) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	log.Info("FailureExample.Hello接口收到请求，返回错误")
	if time.Now().Second()%3 == 0 {
		return errors.New("some_id", "some_biz_detail", 1001)
	}

	return errors.New("some_id", "some_detail", 999)
}

func main() {
	service := stack.NewService(
		stack.Name("stack.rpc.greeter.retry"),
	)

	service.Init()

	// 注册 example handler
	proto.RegisterGreeterHandler(service.Server(), new(FailureExample))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
