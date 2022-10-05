package grpc

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	chatpb "github.com/yiaw/grpc-example/protos/v2/chat"
	userpb "github.com/yiaw/grpc-example/protos/v2/user"
)

func NewGateWay(port int) (http.Handler, error) {
	conn, err := grpc.DialContext(
		context.Background(),
		fmt.Sprintf("0.0.0.0:%d", port),
		grpc.WithBlock(),
		grpc.WithInsecure(),
	)

	if err != nil {
		log.Printf("DialContext is Fail..")
		return nil, err
	}

	mux := runtime.NewServeMux(
		runtime.WithIncomingHeaderMatcher(custom),
	)

	for _, f := range []func(context.Context, *runtime.ServeMux, *grpc.ClientConn) error{
		userpb.RegisterUserHandler,
		chatpb.RegisterChatServiceHandler,
	} {
		if err := f(context.Background(), mux, conn); err != nil {
			log.Printf("Registry Handler Fail.. err=%s", err.Error())
			return nil, err
		}
	}
	return mux, nil
}

func custom(key string) (string, bool) {
	fmt.Printf("input key = %s\n", key)
	switch key {
	case "Token":
		return key, true
	case "accesstoken":
		return key, true
	case "Test":
		return key, true
	default:
		return runtime.DefaultHeaderMatcher(key)
	}

}
