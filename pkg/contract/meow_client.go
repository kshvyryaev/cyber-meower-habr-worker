package contract

type MeowClient interface {
	Create(body string) error
}
