package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func MapBodyTo[T interface{}](resp *http.Response, to *T) error {
	if (resp.StatusCode != http.StatusOK) {
		return fmt.Errorf("error fetching slack users: %s", resp.Status)
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.Unmarshal(data, to)
}


