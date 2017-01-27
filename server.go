package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mateuszdyminski/distlog/logs"
	"github.com/mateuszdyminski/distlog/service"
	"github.com/pborman/uuid"
	"github.com/uber-go/zap"
	"google.golang.org/grpc"
)

var (
	host     = flag.String("host", "localhost", "Host")
	httpPort = flag.Int("http-port", 8080, "Http port")
	rpcHost  = flag.String("rpc-host", "localhost", "Host of RPC server")
	rpcPort  = flag.Int("rpc-port", 8070, "RPC server port")
)

type calculator struct {
	rpcConn   *grpc.ClientConn
	rpcClient service.CalculatorClient
}

func main() {
	flag.Parse()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", *rpcHost, *rpcPort), grpc.WithInsecure())
	if err != nil {
		log.Fatal("can't connect to gRPC server", err)
	}

	calc := &calculator{rpcConn: conn, rpcClient: service.NewCalculatorClient(conn)}

	http.Handle("/calculate", injectCtx(calc.calculate))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *httpPort), nil); err != nil {
		log.Fatalf("can't start http server. err: %v", err)
	}
}

func (calc *calculator) calculate(resp http.ResponseWriter, req *http.Request, rqCtx context.Context) {
	arg1 := req.URL.Query().Get("arg1")
	arg2 := req.URL.Query().Get("arg2")
	operation := req.URL.Query().Get("operation")

	result, err := calc.calculateRemotely(rqCtx, arg1, arg2, operation)
	if err != nil {
		http.Error(resp, fmt.Sprintf("can't calculate values: %s, %s operation: %s! err: %v", arg1, arg2, operation, err), http.StatusBadRequest)
		return
	}

	fmt.Fprint(resp, result)
}

func (calc *calculator) calculateRemotely(ctx context.Context, arg1, arg2, operation string) (int32, error) {
	op, found := service.Operation_value[operation]
	if !found {
		return 0, fmt.Errorf("operation '%d' not supported", operation)
	}

	arg1Int, err := strconv.Atoi(arg1)
	if err != nil {
		return 0, fmt.Errorf("can't parse arg1 '%s', err: %v", arg1, err)
	}

	arg2Int, err := strconv.Atoi(arg2)
	if err != nil {
		return 0, fmt.Errorf("can't parse arg2 '%s', err: %v", arg2, err)
	}

	logs.Logger(ctx).Info("got arguments from http request", zap.Int("arg1", arg1Int), zap.Int("arg2", arg2Int), zap.String("operation", operation))

	result, err := calc.rpcClient.Calculate(ctx, &service.Input{Arg1: int32(arg1Int), Arg2: int32(arg2Int), Operation: service.Operation(op)})
	if err != nil {
		return 0, err
	}

	logs.Logger(ctx).Info("got result for arguments", zap.Int("arg1", arg1Int), zap.Int("arg2", arg2Int), zap.String("operation", operation), zap.Int("result", int(result.Result)))

	return result.Result, nil
}

var httpContext = context.Background()

// injectCtx injects net/context to each request.
func injectCtx(endpointHandler func(http.ResponseWriter, *http.Request, context.Context)) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		endpointHandler(res, req, logs.WithRqId(httpContext, uuid.NewRandom().String()))
	})
}
