package worker

import (
	"sync"

	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/contract"
	"go.uber.org/zap"
)

type HabrUploaderWorker struct {
	meowClient contract.MeowClient
	logger     *zap.Logger
}

func ProvideHabrUploaderWorker(meowClient contract.MeowClient, logger *zap.Logger) *HabrUploaderWorker {
	return &HabrUploaderWorker{
		meowClient: meowClient,
		logger:     logger,
	}
}

func (worker *HabrUploaderWorker) Run(channel chan []string, wg *sync.WaitGroup) {
	defer wg.Done()

	for titles := range channel {
		worker.logger.Info("uploading best titles started")

		for _, title := range titles {
			err := worker.meowClient.Create(title)
			if err != nil {
				worker.logger.Error("title didn't create", zap.String("body", title), zap.Error(err))
			}
			worker.logger.Info("title created", zap.String("body", title))
		}

		worker.logger.Info("uploading best titles finished")
	}
}
