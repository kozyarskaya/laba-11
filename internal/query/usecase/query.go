package usecase

func (u *Usecase) FetchQuery() (string, error) {
	nam, err := u.p.SelectName()
	msg := "Hello," + nam + "!"
	if err != nil {
		return "", err
	}

	if msg == "" {
		return u.defaultMsg, nil
	}

	return msg, nil
}

func (u *Usecase) SetQuery(name string) error {
	err := u.p.InsertQuery(name)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) ChangeQuery(name string) error {
	err := u.p.UpdateQuery(name)
	if err != nil {
		return err
	}

	return nil
}
