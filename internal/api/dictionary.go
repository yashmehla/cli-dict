package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yashmehla/cli-dict/internal/models"
)

func Lookup(word string) (*models.Response, error) {

	url := "https://api.dictionaryapi.dev/api/v2/entries/en/" + word

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("word not found")
	}

	var data models.Response

	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
