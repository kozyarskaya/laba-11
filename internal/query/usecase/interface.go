package usecase

type Provider interface {
	SelectName() (string, error)
	InsertQuery(string) error
	UpdateQuery(string) error
}
