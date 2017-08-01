package feedback

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type FeedbackCommand struct {
	RetroId string     `short:"r" long:"retro-id" description:"Retro Board Id"`
	Meh     MehCommand `command:"meh" alias:"m" description:"Raise a potential concern"`
}

type MehCommand struct {
	Description string `short:"d" long:"description" description:"Write your message" required:"true"`
}

type Category string

const (
	CategoryMeh Category = "meh"
)

type RetroItem struct {
	Description string   `json:"description"`
	Category    Category `json:"category"`
}

var FeedBack FeedbackCommand

func (m *MehCommand) Execute(args []string) error {
	r := RetroItem{
		Description: m.Description,
		Category:    CategoryMeh,
	}
	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://retro-api.cfapps.io/retros/330/items", bytes.NewBuffer(b))
	if err != nil {
		return errors.New(fmt.Sprintf("Error building request\n%v", err.Error()))
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	fmt.Println(req.Body)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.Status != "200" {
		b, _ := httputil.DumpResponse(resp, true)
		return fmt.Errorf("unexpected response code (%d) - %s", resp.StatusCode, string(b))
	}
	return nil
}
