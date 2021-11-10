package client

import (
	"context"
	"time"

	"github.com/kshvyryaev/cyber-meower-proto/pkg/proto"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

type GrpcMeowClient struct {
	client proto.MeowServiceClient
}

func ProvideGrpcMeowClient(meowerServiceConnection *grpc.ClientConn) *GrpcMeowClient {
	client := proto.NewMeowServiceClient(meowerServiceConnection)

	return &GrpcMeowClient{
		client: client,
	}
}

func (client *GrpcMeowClient) Create(body string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	request := &proto.CreateMeowRequest{
		Body: body,
	}

	_, err := client.client.Create(ctx, request)
	if err != nil {
		return errors.Wrap(err, "grpc meow client")
	}

	return nil
}
