package grpc

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
)

func Auth() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		fmt.Printf(">>>>>>>> AUTH CALL\n")
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			v := md["token"]
			if len(v) == 0 {
				return nil, status.Errorf(codes.PermissionDenied, "not found token")
			}

			// add metadata
			md.Set("aauth", "ok")
			md.Set("addmetadata", "ok")

			newCtx := metadata.NewIncomingContext(ctx, md)

			return handler(newCtx, req)
		}
		return nil, status.Errorf(codes.PermissionDenied, "not found incomming context")
	}
}

func Audit() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		fmt.Printf(">>>>>> AUDIT CALL\n")
		startTime := time.Now()

		// call handler
		resp, err := handler(ctx, req)

		endTime := time.Since(startTime)
		elapsedTime := float32(endTime) / float32(1000)

		elapsed := fmt.Sprintf("%0.3fuq", elapsedTime)
		fmt.Printf("<<<<<< AUDIT END :: START_TIME=[%s] ELAPSED_TIME=[%s]\n", startTime.Format("2006-01-02 15:04:05.000000"), elapsed)

		return resp, err

	}
}

func Log() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		fmt.Printf(">>>> LOG CALL\n")
		var (
			code     string
			message  string
			peerAddr string
			auth     string
			metaadd  string
		)
		// handler call
		resp, err := handler(ctx, req)
		pr, ok := peer.FromContext(ctx)
		if ok {
			peerAddr = pr.Addr.String()
		}

		errStatus, ok := status.FromError(err)
		if ok {
			code = errStatus.Code().String()
			message = errStatus.Message()
		} else {
			code = codes.Unknown.String()
			message = "unknown code"
		}

		if md, ok := metadata.FromIncomingContext(ctx); ok {
			tempAuth := md.Get("aauth")
			if len(tempAuth) != 0 {
				auth = tempAuth[0]
			} else {
				auth = "no"
			}
			tempAdd := md.Get("addmetadata")
			if len(tempAdd) != 0 {
				metaadd = tempAdd[0]
			} else {
				metaadd = "no"
			}
		}

		fmt.Printf("<<<< LOG END :: GRPC UNARY CALL METHOD=[%s] CODE=[%s] MESSAGE=[%s] PEER_ADDRESS=[%s] AUTH_OK=[%s] ADD_META=[%s]\n",
			info.FullMethod, code, message, peerAddr, auth, metaadd)
		return resp, err

	}
}

func customMiddleware() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			data, _ := json.MarshalIndent(md, "", " ")
			log.Printf("Request Incomming Header\n")
			log.Printf("%s\n", string(data))
		}

		if md, ok := metadata.FromOutgoingContext(ctx); ok {
			data, _ := json.MarshalIndent(md, "", " ")
			log.Printf("Request Outgoing Header\n")
			log.Printf("%s\n", string(data))

		}

		log.Printf("context %#v\n", ctx)

		resp, err := handler(ctx, req)

		return resp, err
	}
}

func NewStreamMiddleware() grpc.StreamServerInterceptor {
	logrus.ErrorKey = "grpc.error"
	logrusEntry := logrus.NewEntry(logrus.StandardLogger())
	middleware := grpc_middleware.ChainStreamServer(
		grpc_ctxtags.StreamServerInterceptor(),
		grpc_logrus.StreamServerInterceptor(logrusEntry),
		grpc_recovery.StreamServerInterceptor(),
	)
	return middleware
}

func NewChainMiddleware() grpc.UnaryServerInterceptor {
	middleware := grpc_middleware.ChainUnaryServer(
		grpc_ctxtags.UnaryServerInterceptor(),
		Audit(),
		Log(),
		Auth(),
		grpc_recovery.UnaryServerInterceptor(),
	)

	return middleware
}
