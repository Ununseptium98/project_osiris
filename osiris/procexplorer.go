package osiris

import (
	"encoding/json"
	"log"

	"github.com/shirou/gopsutil/process"
)

func GetProcessExePath(PID int32) (string, error) {
	/*
		Takes a PID as input, returns the path to its exe
	*/

	proc, err := process.NewProcess(PID)

	if err != nil { //error with the PID, couldn'r get the process associated

		return "null", err

	}

	path, err := proc.Exe()

	if err != nil { //return the error if couldn't get the exe

		return "null", err

	}
	return path, err

}

func GetProcessExeHashMap() map[string]string {
	//creates a map of exe path and the hash of the exe

	pidProcessMap := make(map[string]string)

	pid_table, err := process.Pids() //retrieve table of PID of running processes
	if err != nil {
		log.Fatal(err)
	}

	for _, pid := range pid_table { //Creates a map with exe path with exe hash

		processPath, err := GetProcessExePath(pid)

		if err != nil || processPath == "null" { //If there was an error reading from the PID, pass it

			continue
		}
		//associate each prpcess exe path to its hash in a map
		pidProcessMap[processPath] = PathTomd5(processPath)

	}

	return pidProcessMap

}

func GetProcessExeHashJson() []byte {
	/*
		Returns process exe hash with JSON format
	*/

	exeHashMap := GetProcessExeHashMap()

	json, err := json.Marshal(exeHashMap)

	if err != nil {
		log.Fatal(err)
	}

	return json

}
