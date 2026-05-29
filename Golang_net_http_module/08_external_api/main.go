package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

type catResponce struct {
	Fact   string `json:"fact"`
	Length int    `json:"length"`
}

func writeJson(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	_ = json.NewEncoder(w).Encode(data)
}

func fetchResponse() (catResponce, error) {
	url := "https://catfact.ninja/fact"

	res, err := http.Get(url)
	if err != nil {
		return catResponce{}, err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return catResponce{}, fmt.Errorf("external api failes: %s", res.Status)
	}

	var data catResponce

	bodyByte, err := io.ReadAll(res.Body)
	if err != nil {
		return catResponce{}, err
	}

	if err := json.Unmarshal(bodyByte, &data); err != nil {
		return catResponce{}, err
	}

	return data, nil

}

func externalHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		writeJson(w, http.StatusMethodNotAllowed, map[string]any{
			"ok":    false,
			"error": "only get method allowed",
		})
		return
	}

	data, err := fetchResponse()
	if err != nil {
		writeJson(w, http.StatusBadGateway, map[string]any{
			"ok":    false,
			"error": "failed to fetch data",
		})
		return
	}

	writeJson(w, http.StatusOK, map[string]any{
		"ok":   true,
		"time": time.Now().UTC(),
		"external": map[string]any{
			"source": "cat.ninja",
			"fact":   data.Fact,
			"length": data.Length,
		},
	})

}

func main() {

	http.HandleFunc("/external", externalHandler)

	err := http.ListenAndServe(":5000", nil)
	fmt.Println(err)
}
