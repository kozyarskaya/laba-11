package usecase

type Provider interface {
	SelectCount() (string, error)
	InsertCount(float32) error
	UpdateCount(float32) error
}
