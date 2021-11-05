package client

type MeowClient interface {
	Create(body string) error
}
