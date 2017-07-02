package utils

import (
	"net/http"
	"bytes"
)

func HttpPostJson(endpoint, username, password, payload string) error {
	var jsonStr = []byte(payload)
	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonStr))
	req.SetBasicAuth(username, password)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
			panic(err)
	}

	defer resp.Body.Close()
	return err
}
