package osiris

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type Request struct {
	RequestID             string   `json:"RequestID"`
	ProcessTree           bool     `json:"ProcessTree"`
	HashingProcessTree    bool     `json:"HashingProcessTree"`
	DumpRegistryKey       bool     `json:"DumpRegistryKey"`
	Keys                  []string `json:"Keys"`
	WatchListAppend       bool     `json:"WatchListAppend"`
	PathsToWatch          []string `json:"PathsToWatch"`
	CheckWatchListChanges bool     `json:"CheckWatchListChanges"`
}

func ReadTaskJson(taskJSON string) Request {
	byteValue, err := ioutil.ReadFile(taskJSON)

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

func ExecuteTaskJson(taskJSON string) {

	req := ReadTaskJson(taskJSON)

	requestID := req.RequestID

	if req.ProcessTree {
		WriteProcessPathJson(requestID)
	}

	if req.HashingProcessTree {
		WriteProcessExeHashJson(req.RequestID)
	}

	if req.DumpRegistryKey {

		for iterator, key := range req.Keys {
			GetAllKeyValuesJSON(key, 1, true,
				req.RequestID+"_key"+strconv.FormatInt(int64(iterator), 10))

		}
	}

	if req.WatchListAppend {
		CreateWatchList()

		for _, path := range req.PathsToWatch {
			AppendWatchList(path)
		}
	}

}
