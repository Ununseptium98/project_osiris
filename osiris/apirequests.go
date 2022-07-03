package osiris

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"os/user"
)

func EnrollAgent(host string) {

	hostname, _ := os.Hostname()
	currentUser, _ := user.Current()
	username := currentUser.Name

	//Encode the data to JSON for the API request
	postBody, _ := json.Marshal(map[string]string{
		"pc_name":  hostname,
		"username": username,
	})

	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(host+"/client/"+getMacAddr(), "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//read the response

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	stringBody := string(body)
	log.Printf(stringBody)

}

type dataType struct {
	Hash_id   string `json:"hash_id"`
	Path_file string `json:"path_file"`
}

type hashReq struct {
	Mac_addr  string   `json:"mac_addr"`
	Data_type string   `json:"data_type"`
	Data      dataType `json:"data"`
}

func SendHash(host string, filePath string) {

	//Creates the data for the request

	data := dataType{
		Hash_id:   PathTomd5(filePath),
		Path_file: filePath,
	}
	req := &hashReq{
		Mac_addr:  getMacAddr(),
		Data_type: "HASH_FILE",
		Data:      data,
	}

	//transforms it to JSON
	postBody, err := json.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(postBody))
	//
	responseBody := bytes.NewBuffer(postBody)

	resp, err := http.Post(host+"/client/"+getMacAddr(), "application/json", responseBody)

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	//read the response

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatal(err)
	}

	stringBody := string(body)
	log.Printf(stringBody)

}

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}