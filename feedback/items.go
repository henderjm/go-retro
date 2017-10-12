package feedback

import (
	"fmt"

	gotabulate "github.com/bndr/gotabulate"
	table "github.com/henderjm/go-feedback/table"
)

type ItemsCommand struct {
}

func (a *ItemsCommand) Execute(args []string) error {
	url := fmt.Sprintf("https://retro-api.cfapps.io/retros/%v/", FeedBack.RetroId)

	board, err := GetRetroBoard(url)
	if err != nil {
		return err
	}
	var happy []string
	var meh []string
	var sad []string

	for _, a := range board.Board.RetroItems {
		if a.Category == CategoryHappy {
			happy = append(happy, a.Description)
		}

		if a.Category == CategoryMeh {
			meh = append(meh, a.Description)
		}

		if a.Category == CategorySad {
			sad = append(sad, a.Description)
		}
	}

	outputTable(happy, meh, sad)

	return nil
}

func outputTable(h, m, s []string) error {
	contents, err := table.PadRowsToColumns(h, m, s, "")
	if err != nil {
		return err
	}

	t := gotabulate.Create(contents)
	t.SetHeaders([]string{"Happy", "Meh", "Sad"})

	return nil
}
