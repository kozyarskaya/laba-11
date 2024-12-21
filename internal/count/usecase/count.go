package usecase

func (u *Usecase) FetchCount() (string, error) {
	msg, err := u.p.SelectCount()
	if err != nil {
		return "", err
	}
	return msg, nil
}

func (u *Usecase) SetCount(v float32) error {
	err := u.p.InsertCount(v)
	if err != nil {
		return err
	}

	return nil
}

func (u *Usecase) ChangeCount(v float32) error {
	err := u.p.UpdateCount(v)
	if err != nil {
		return err
	}

	return nil
}
