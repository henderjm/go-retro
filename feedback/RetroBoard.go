package feedback

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
)

func GetRetroBoard(url string) (*RetroBoard, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		b, _ := httputil.DumpResponse(resp, true)
		return nil, fmt.Errorf("unexpected response code (%d) - %s", resp.StatusCode, string(b))
	}

	b := new(bytes.Buffer)
	_, err = b.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	var actions = new(RetroBoard)
	err = json.Unmarshal(b.Bytes(), &actions)
	if err != nil {
		return nil, err
	}

	return actions, nil
}
