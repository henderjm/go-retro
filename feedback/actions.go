package feedback

import (
	"fmt"
)

type ActionsCommand struct {
}

func (a *ActionsCommand) Execute(args []string) error {
	url := fmt.Sprintf("https://retro-api.cfapps.io/retros/%v/", FeedBack.RetroId)

	board, err := GetRetroBoard(url)
	if err != nil {
		return err
	}

	for _, a := range board.Board.ActionItems {
		fmt.Println(a.Description)
	}

	return nil
}
