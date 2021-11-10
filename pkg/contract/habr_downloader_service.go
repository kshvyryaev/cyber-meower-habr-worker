package contract

type HabrDownloaderService interface {
	DownloadBestTitles() ([]string, error)
}
