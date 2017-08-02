package main

import (
	"os"

	"fmt"

	"github.com/henderjm/go-feedback/feedback"
	"github.com/jessevdk/go-flags"
)

func main() {
	p := flags.NewParser(&feedback.FeedBack, flags.Default)
	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			fmt.Println("What is wrong")
			os.Exit(0)
		} else {
			os.Exit(0)
		}
	}
	fmt.Println("No errors")
	os.Exit(0)
}
