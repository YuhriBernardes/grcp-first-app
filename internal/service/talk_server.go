package service

import (
	"context"
	"errors"
	"log"
	"net"
	"strconv"

	"github.com/YuhriBernardes/grpc-first-app/internal/grpc/talk"
	"google.golang.org/grpc"
)

type TalkServiceServer struct {}

func (t *TalkServiceServer) Ask (ctx context.Context, request *talk.ServiceQuestion) (response *talk.ServiceAnswer, err error){
	question := request.GetQuestion();

	if question == ""{
		return &talk.ServiceAnswer{}, errors.New("Question can't be empty")
	}

	answer := "Q:" + question + "\nA: answer from server "+t.ServiceName


	return &talk.ServiceAnswer{Response: answer}, nil
}

func Init (serverName string, host string, port int){
	address := host + ":" + strconv.Itoa(port)
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to create listener %v", err)
	}



	grpcServer := grpc.NewServer()
	talk.RegisterServiceTalkService(grpcServer, &TalkServiceServer{})

	if err := grpcServer.Serve(lis); err !=nil{
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}

}
