package groupie

import (
	"encoding/json"
	"net/http"
)

func fetch(url string, endpoint string, data interface{}) error {
	res, err := http.Get(url + endpoint)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	
	json.NewDecoder(res.Body).Decode(data)
	return nil
}
