package feedback

type StartRetroCommand struct {
}

func (s *StartRetroCommand) Execute(args []string) error {
	url := "https://retro-api.cfapps.io/retros/330/"

	_, err := GetRetroBoard(url)
	if err != nil {
		return err
	}

	return nil
}

func (i *RetroItem) MarkItemAsDone() error {
	i.Done = true
	return nil
}
