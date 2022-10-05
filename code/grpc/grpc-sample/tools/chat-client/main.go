package main

import (
	"context"
	"fmt"
	"log"

	chatpb "github.com/yiaw/grpc-example/protos/v2/chat"
	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("0.0.0.0:8090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("connect fail %v\n", err)
	}
	defer conn.Close()

	cli := chatpb.NewChatServiceClient(conn)

	var u string
	fmt.Printf("enter username : ")
	fmt.Scanf("%s", &u)

	stream, err := cli.RegistRouterChannel(context.Background())
	if err != nil {
		log.Fatalf("RegistRouterChannel %v\n", err)
	}

	go func() {
		for {
			msg, err := stream.Recv()
			if err != nil {
				log.Printf("stream message recv fail err %v\n", err)
				return
			}
			fmt.Printf("[%s] %s\n", msg.User, msg.Msg)
		}
	}()

	for {
		var msg string
		fmt.Scanf("%s", &msg)
		m := &chatpb.Message{
			User: u,
			Msg:  msg,
		}
		_, err := cli.SendMessage(context.Background(), m)
		if err != nil {
			log.Printf("Send Message Fail.. err %v\n", err)
			continue
		}
	}
}
