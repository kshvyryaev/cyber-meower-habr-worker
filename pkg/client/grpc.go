package client

import (
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
)

func ProvideMeowerServiceGrpcConnection(config *pkg.Config) (*grpc.ClientConn, func(), error) {
	connection, err := grpc.Dial(config.MeowerServiceAddress, grpc.WithInsecure())
	if err != nil {
		return nil, nil, errors.Wrap(err, "meower service grpc connection")
	}

	cleanup := func() {
		connection.Close()
	}

	return connection, cleanup, nil
}
