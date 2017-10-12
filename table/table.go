package table

import (
	"bytes"
	"fmt"
)

type Table struct {
	Default_Empty string ""
	Format        tableformat
}

type tableformat string

const (
	Padded tableformat = "padded"
	Simple tableformat = "simple"
)

func (t *Table) Render(d ...[]string) (string, error) {
	var result []string

	for _, x := range d {
		for _, y := range x {
			result = append(result, y)
		}
	}

	var buf bytes.Buffer

	for _, r := range result {
		if t.Format == Padded {
			buildPaddedColumn(r)
		} else if t.Format == Simple {
			buf.WriteString(fmt.Sprintf("%v\n", r))
		}
	}

	return buf.String(), nil
}

func determineColumnMeta(d ...[]string) (int, int) {
	maxRows := determineMaxRows(d[0])
	maxColWidth := determineColumnWidth(d[0]) // TODO Let's tidy this up :)

	return maxRows, maxColWidth
}

func determineMaxRows(d ...[]string) int {
	return len(d)
}

func determineColumnWidth(columnData []string) int {
	maxCol := 0
	for _, x := range columnData {
		if len(x) > maxCol {
			maxCol = len(x)
		}
	}
	return maxCol
}

func buildPaddedColumn(value string) {

}
