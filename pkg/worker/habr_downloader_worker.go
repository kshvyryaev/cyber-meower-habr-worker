package worker

import (
	"sync"
	"time"

	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/service"
	"go.uber.org/zap"
)

type HabrDownloaderWorker struct {
	habrDownloaderService *service.HabrDownloaderService
	logger                *zap.Logger
	channel               chan []string
}

func ProvideHabrDownloaderWorker(
	habrDownloaderService *service.HabrDownloaderService,
	logger *zap.Logger) (*HabrDownloaderWorker, func()) {
	channel := make(chan []string)

	cleanup := func() {
		close(channel)
	}

	return &HabrDownloaderWorker{
		habrDownloaderService: habrDownloaderService,
		logger:                logger,
		channel:               channel,
	}, cleanup
}

func (worker *HabrDownloaderWorker) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	// TODO: Move it to configuration
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		worker.logger.Info("Downloading best titles started")
		titles, err := worker.habrDownloaderService.DownloadBestTitles()
		if err != nil {
			worker.logger.Error("Downloading best titles finished with errors: ", zap.Error(err))
		}
		worker.logger.Info("Downloading best titles finished succesfuly")

		worker.channel <- titles
		worker.logger.Info("Best titles sent to channel")
	}
}

func (worker *HabrDownloaderWorker) GetChannel() chan []string {
	return worker.channel
}
