package feedback

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

type ActionsCommand struct{}

type Item struct {
	Description int64 `json:"id"`
}

type ActionList struct {
	ActionItems Item `json:"retro"`
}

func (a *ActionsCommand) Execute(args []string) error {
	url := "https://retro-api.cfapps.io/retros/330/"

	fmt.Println("Actions")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		b, _ := httputil.DumpResponse(resp, true)
		return fmt.Errorf("unexpected response code (%d) - %s", resp.StatusCode, string(b))
	}

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return err
	}

	var actions = new(ActionList)
	err = json.Unmarshal(b.Bytes(), &actions)
	if err != nil {
		return err
	}

	fmt.Println(actions)
	//for _, a := range actions.ActionItems {
	//	fmt.Println(a.Description)
	//}

	return nil
}
