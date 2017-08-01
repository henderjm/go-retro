package feedback

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type HappyCommand struct {
	Description string `short:"d" long:"description" description:"Write your message" required:"true"`
}

func (h *HappyCommand) Execute(args []string) error {
	r := RetroItem{
		Description: h.Description,
		Category:    CategoryHappy,
	}

	b, err := json.Marshal(r)
	if err != nil {
		return err
	}

	url := fmt.Sprintf("https://retro-api.cfapps.io/retros/%v/items", FeedBack.RetroId)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(b))
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
