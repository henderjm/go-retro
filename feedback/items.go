package feedback

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
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
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Happy", "Meh", "Sad"})

	items := [][]string(
	for _, v := range h {
		table.Append(v)
	}
	return nil
}
