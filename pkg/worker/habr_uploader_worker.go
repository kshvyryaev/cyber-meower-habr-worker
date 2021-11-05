package worker

import (
	"sync"

	"go.uber.org/zap"
)

type HabrUploaderWorker struct {
	logger *zap.Logger
}

func ProvideHabrUploaderWorker(logger *zap.Logger) *HabrUploaderWorker {
	return &HabrUploaderWorker{
		logger: logger,
	}
}

func (worker *HabrUploaderWorker) Run(channel chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for range channel {
		worker.logger.Info("Uploading best titles started")
		// TODO: Implement it
		worker.logger.Info("Uploading best titles finished")
	}
}
