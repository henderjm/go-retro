package feedback

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os/user"

	"golang.org/x/crypto/ssh/terminal"
)

type LoginCommand struct {
}

type retro struct {
	Password string `json:"password"`
}

type loginParams struct {
	Retro retro `json:"retro"`
}

type loginResponse struct {
	Token string `json:"token"`
}

func (l *LoginCommand) Execute(args []string) error {
	fmt.Println("Enter retro password: ")
	password, err := terminal.ReadPassword(0)
	if err != nil {
		return err
	}
	r := retro{
		Password: string(password),
	}

	params := loginParams{
		Retro: r,
	}

	url := fmt.Sprintf("https://retro-api.cfapps.io/retros/%v/login", FeedBack.RetroId)

	b, err := json.Marshal(params)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(b))
	if err != nil {
		return errors.New(fmt.Sprintf("Error building request\n%v", err.Error()))
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

	usr, err := user.Current()
	if err != nil {
		return err
	}

	var lResponse = new(loginResponse)

	serverResponse, err := ioutil.ReadAll(resp.Body)
	json.Unmarshal(serverResponse, lResponse)
	if err != nil {
		return err
	}

	bearerToken := fmt.Sprintf(`{"token": "Bearer %v"}`, lResponse.Token)

	err = ioutil.WriteFile(fmt.Sprintf("%v/feedback-config.json", usr.HomeDir), []byte(bearerToken), 0644)
	return nil
}
