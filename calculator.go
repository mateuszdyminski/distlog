package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/mateuszdyminski/distlog/logs"
	"github.com/mateuszdyminski/distlog/service"
	"github.com/uber-go/zap"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	rpcHost = flag.String("rpc-host", "localhost", "Host")
	rpcPort = flag.Int("rpc-port", 8070, "RPC port")
)

type calculatorServer struct{}

func main() {
	flag.Parse()

	srv := &calculatorServer{}
	err := srv.startGrpc()
	if err != nil {
		log.Fatalf("can't start rpc server. err: %v", err)
	}
}

func (s *calculatorServer) startGrpc() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *rpcHost, *rpcPort))
	if err != nil {
		return err
	}

	server := grpc.NewServer(grpc.MaxMsgSize(1024 * 1024 * 20)) // 20mb
	service.RegisterCalculatorServer(server, s)

	if err := server.Serve(lis); err != nil {
		return err
	}

	return nil
}

func (s *calculatorServer) Calculate(ctx context.Context, input *service.Input) (*service.Result, error) {
	logs.Logger(ctx).Info("got arguments from rpc request", zap.Int("arg1", int(input.Arg1)), zap.Int("arg2", int(input.Arg2)), zap.String("operation", input.Operation.String()))

	switch input.Operation {
	case service.Operation_ADD:
		return &service.Result{Result: input.Arg1 + input.Arg2}, nil
	case service.Operation_SUB:
		return &service.Result{Result: input.Arg1 - input.Arg2}, nil
	case service.Operation_MULTI:
		return &service.Result{Result: input.Arg1 * input.Arg2}, nil
	case service.Operation_DIVIDE:
		if input.Arg2 == 0 {
			return nil, errors.New("can't divide by zero!")
		}
		return &service.Result{Result: input.Arg1 / input.Arg2}, nil
	default:
		return nil, errors.New("unknow operation")
	}
}
