package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type APIResponse[T any] struct {
	Data []T         `json:"data"`
	Meta interface{} `json:"meta"`
}

var httpClient = &http.Client{
	Timeout: 10 * time.Second,
}

func FetchJSON[T any](req *http.Request, res APIResponse[T], log *zap.Logger) ([]T, error) {
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Error("Error request fetch", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed fetch %s: status %d", req.URL, resp.StatusCode)
	}

	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		log.Error("Error decode response", zap.Error(err))
		return nil, err
	}

	return res.Data, nil
}
