package talariaClientDemo

import (
	"context"
	"encoding/json"
	talaria "github.com/kelindar/talaria/client/golang"
	"log"
	"testing"
	"time"
)

func TestTalariaClient(t *testing.T) {
	var timeout time.Duration = 5
	var concurrency int = 10
	var errorPercentage int = 50
	c, err := New("www.talaria.net:8043", &timeout, &concurrency, &errorPercentage)

	if err != nil {
		log.Fatal("Error is " + err.Error())
	}

	testTime := time.Now()
	testEvents := []talaria.Event{
		{
			"event": "abc",
		},
		{
			"value": 123,
		},
		{
			"mutable": true,
		},
		{
			"time": testTime,
		},
		{
			"misc": json.RawMessage("{1:2}"),
		},
		{
			"event": "xyz",
			"time":  456,
			"color": "yellow",
			"topic": "movie",
		},
		{
			"event": "abc",
			"nested": map[string]interface{}{
				"level0": 0,
				"level1": map[string]interface{}{
					"test": 1,
					"level2": []int{
						1,
						2,
						3,
					},
				},
			},
		},
	}

	c.IngestBatch(context.Background(), testEvents)
}

