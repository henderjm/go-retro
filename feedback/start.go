package feedback

import (
	"fmt"
	"log"
	"time"
)

type StartRetroCommand struct {
}

func (s *StartRetroCommand) Execute(args []string) error {
	url := "https://retro-api.cfapps.io/retros/330/"

	b, err := GetRetroBoard(url)
	if err != nil {
		return err
	}
	LetsRetro(b)
	return nil
}

func LetsRetro(rb *RetroBoard) {
	for _, v := range rb.Board.RetroItems {
		err := StartItem(rb.Board.Slug, v)
		if err != nil {
			panic(err)
		}
		for startTime(v) {
			fmt.Println("You're not done????")
		}
		v.MarkItemAsDone()
		Patch(rb.Board.Slug, v)
	}
}

func startTime(i RetroItem) bool {
	fmt.Println(fmt.Sprintf("%v: %v", i.Category, i.Description))
	timeChan := time.NewTimer(time.Second * 10).C
	userChan := make(chan bool)

	fmt.Println("Finished?: (y, Y, n, N)")
	go areWeDone(userChan)
	for {
		select {
		case <-timeChan:
			return false
		case x := <-userChan:
			return x
		}
	}
}

func areWeDone(c chan bool) {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil {
		log.Fatal(err)
	}
	okayResponses := []string{"y", "Y"}
	nokayResponses := []string{"n", "N"}
	if containsString(okayResponses, response) {
		c <- true
	} else if containsString(nokayResponses, response) {
		c <- false
	} else {
		fmt.Println("Please type yes or no and then press enter:")
		areWeDone(c)
	}
}

func containsString(slice []string, element string) bool {
	return !(posString(slice, element) == -1)
}

func posString(slice []string, element string) int {
	for index, elem := range slice {
		if elem == element {
			return index
		}
	}
	return -1
}
