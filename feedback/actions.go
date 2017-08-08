package feedback

import (
	"fmt"
)

type ActionsCommand struct {
}

type Item struct {
	Description int64 `json:"id"`
}

func (a *ActionsCommand) Execute(args []string) error {
	url := "https://retro-api.cfapps.io/retros/330/"

	board, err := GetRetroBoard(url)
	if err != nil {
		return err
	}

	for _, a := range board.Board.ActionItems {
		fmt.Println(a.Description)
	}

	return nil
}
