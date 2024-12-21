package api

type Usecase interface {
	FetchQuery() (string, error)
	SetQuery(name string) error
	ChangeQuery(name string) error
}
