package main

import (
	"encoding/json"
	"sync/atomic"

	vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
	tr := func(id uint64) vegeta.Targeter {
		type entity struct {
			Name string `json:"entityName"`
			ID   uint64 `json:"entityId"`
		}
		return func(t *vegeta.Target) (err error) {
			t.Method = "POST"
			t.URL = "http://localhost:6060"

			t.Body, err = json.Marshal(&entity{
				Name: "burger",
				ID:   atomic.AddUint64(&id, 1),
			})

			return err
		}
	}(0)
	_ = tr

	/*
		jq -ncM 'while(true; .+1) | {method: "POST", url: "http://localhost:6060", body: {entityName: "burger", entityId: .} | @base64}'  | \
		vegeta attack -lazy --format=json -duration=30s | tee results.bin | vegeta report
	*/

	// Sending requests with dynamic body
	// https://github.com/tsenart/vegeta/issues/330
}
