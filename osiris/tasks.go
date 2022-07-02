package osiris

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type Request struct {
	RequestID             string   `json:"RequestID"`
	ProcessTree           bool     `json:"ProcessTree"`
	HashingProcessTree    bool     `json:"HashingProcessTree"`
	DumpRegistryKey       bool     `json:"DumpRegistryKey"`
	Keys                  []string `json:"Keys"`
	PrintSubkeys          bool     `json:"PrintSubkeys"`
	WatchListAppend       bool     `json:"WatchListAppend"`
	PathsToWatch          []string `json:"PathsToWatch"`
	CheckWatchListChanges bool     `json:"CheckWatchListChanges"`
}

func ReadJson() Request {
	byteValue, err := ioutil.ReadFile("actions.json")

	if err != nil {
		log.Fatal(err)
	}

	var request Request

	json.Unmarshal(byteValue, &request)

	for _, key := range request.Keys {
		fmt.Println("Key ", key)
	}

	return request

}
