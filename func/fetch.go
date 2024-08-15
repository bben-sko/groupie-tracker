package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func fetchAndDecode(url string, target interface{}) error {
	response, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error fetching data: %w", err)
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(target)
	if err != nil {
		return fmt.Errorf("error decoding response: %w", err)
	}

	return nil
}