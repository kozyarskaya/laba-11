package api

type Usecase interface {
	FetchCount() (string, error)
	SetCount(float32) error
	ChangeCount(float32) error
}
