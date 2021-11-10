//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/client"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/contract"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/service"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/worker"
	"go.uber.org/zap"
)

func InitializeHabrDownloaderWorker(logger *zap.Logger) (*worker.HabrDownloaderWorker, func()) {
	panic(wire.Build(
		service.ProvideHabrDownloaderService,
		wire.Bind(new(contract.HabrDownloaderService), new(*service.HabrDownloaderService)),
		worker.ProvideHabrDownloaderWorker,
	))
}

func InitializeHabrUploaderWorker(config *pkg.Config, logger *zap.Logger) (*worker.HabrUploaderWorker, func(), error) {
	panic(wire.Build(
		client.ProvideMeowerServiceGrpcConnection,
		client.ProvideGrpcMeowClient,
		wire.Bind(new(contract.MeowClient), new(*client.GrpcMeowClient)),
		worker.ProvideHabrUploaderWorker,
	))
}
