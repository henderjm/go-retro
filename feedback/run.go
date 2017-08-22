package feedback

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
)

func Run(r RetroItem) error {
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

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		b, _ := httputil.DumpResponse(resp, true)
		return fmt.Errorf("unexpected response code (%d) - %s", resp.StatusCode, string(b))
	}
	return nil
}

func StartItem(slug string, r RetroItem) error {
	id := strconv.FormatInt(r.ID, 10)

	u := fmt.Sprintf("https://retro-api.cfapps.io/retros/%v/discussion", slug)
	resp, err := http.PostForm(u, url.Values{"item_id": {id}})
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	_, err = ioutil.ReadAll(resp.Body)

	if nil != err {
		fmt.Println("errorination happened reading the body", err)
		return err
	}

	return nil
}

func Patch(slug string, r RetroItem) error {
	url := fmt.Sprintf("https://retro-api.cfapps.io/retros/%v/items/%v/done", slug, r.ID)
	req, err := http.NewRequest("PATCH", url, nil)
	if err != nil {
		return errors.New(fmt.Sprintf("Error building request\n%v", err.Error()))
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 201 {
		b, _ := httputil.DumpResponse(resp, true)
		return fmt.Errorf("unexpected response code (%d) - %s", resp.StatusCode, string(b))
	}
	return nil
}
