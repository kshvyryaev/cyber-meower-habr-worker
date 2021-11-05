//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/service"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/worker"
	"go.uber.org/zap"
)

func InitializeHabrDownloaderWorker(logger *zap.Logger) (*worker.HabrDownloaderWorker, func()) {
	panic(wire.Build(
		service.ProvideHabrDownloaderService,
		worker.ProvideHabrDownloaderWorker,
	))
}

func InitializeHabrUploaderWorker(logger *zap.Logger) *worker.HabrUploaderWorker {
	panic(wire.Build(
		worker.ProvideHabrUploaderWorker,
	))
}
