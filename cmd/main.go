package main

import (
	"sync"

	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg"
	"github.com/kshvyryaev/cyber-meower-habr-worker/pkg/di"
)

func main() {
	logger, loggerCleanup, err := pkg.ProvideZap()
	if err != nil {
		panic("cannot initialize zap logger: " + err.Error())
	}
	defer loggerCleanup()

	downloader, downloaderCleanup := di.InitializeHabrDownloaderWorker(logger)
	defer downloaderCleanup()

	uploader := di.InitializeHabrUploaderWorker(logger)

	wg := &sync.WaitGroup{}
	wg.Add(1)
	go downloader.Run(wg)
	wg.Add(1)
	go uploader.Run(downloader.GetChannel(), wg)
	wg.Wait()
}
